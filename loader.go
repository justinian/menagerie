package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/justinian/ark"
	_ "github.com/mattn/go-sqlite3"
)

type Loader struct {
	lock      sync.Mutex
	db        *sqlx.DB
	classMap  *ClassMap
	saveFiles []string
}

func createLoader(dbname string, specfiles, savefiles []string) (*Loader, error) {
	// Always start with a fresh-loaded db, because options could have
	// changed.
	err := os.Remove(dbname)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("Could not move old db file:\n%w", err)
		}
	}

	log.Printf("Creating sqlite3 database: %s", dbname)

	db, err := sqlx.Connect("sqlite3", dbname)
	if err != nil {
		return nil, fmt.Errorf("Could not open db file:\n%w", err)
	}

	for _, table := range databaseSchema {
		_, err = db.Exec(table)
		if err != nil {
			return nil, fmt.Errorf("Could not create SQL schema:\n%w", err)
		}
	}

	classMap, err := readSpecFiles(specfiles...)
	if err != nil {
		return nil, fmt.Errorf("Reading spec files:\n%w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("Could not begin SQL transaction:\n%w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO classes VALUES (?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("Could not prepare SQL class insert:\n%w", err)
	}

	for bpName, class := range classMap.Map {
		_, err = stmt.Exec(class.Id, bpName, class.Name)
		if err != nil {
			return nil, fmt.Errorf("Inserting class: (%d, %s):\n%w", class.Id, class.Name, err)
		}
	}

	log.Printf("Inserted %d class names from %d spec files.", len(classMap.Map), len(specfiles))

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("Could not commit class names:\n%w", err)
	}

	return &Loader{db: db, classMap: classMap, saveFiles: savefiles}, nil
}

func (l *Loader) run() error {
	for _, savefile := range l.saveFiles {
		err := l.processSavefile(savefile)
		if err != nil {
			return fmt.Errorf("Processing %s:\n%w", savefile, err)
		}
	}

	go l.watcher()
	return nil
}

func (l *Loader) processSavefile(filename string) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	log.Printf("Processing save file: %s", filename)

	archive, err := ark.OpenArchive(filename)
	if err != nil {
		return fmt.Errorf("Could not open save file:\n%w", err)
	}

	save, err := ark.ReadSaveGame(archive)
	if err != nil {
		return fmt.Errorf("Could not read save game:\n%w", err)
	}

	worldName := save.DataFiles[0]
	if strings.HasSuffix(worldName, "_P") {
		worldName = worldName[:len(worldName)-2]
	}

	tx, err := l.db.Beginx()
	if err != nil {
		return fmt.Errorf("Could not begin SQL transaction:\n%w", err)
	}

	_, err = tx.Exec(`
		INSERT INTO worlds (name) VALUES (?)
		ON CONFLICT (name) DO UPDATE SET iter=iter+1`, worldName)
	if err != nil {
		return fmt.Errorf("Could not insert world name:\n%w", err)
	}

	var worldId int
	err = tx.Get(&worldId, `SELECT (id) FROM worlds WHERE name = ?`, worldName)
	if err != nil {
		return fmt.Errorf("Could not get world id:\n%w", err)
	}

	_, err = tx.Exec("DELETE FROM dinos WHERE world = ?", worldId)
	if err != nil {
		return fmt.Errorf("Could not clear previous world iteration:\n%w", err)
	}

	err = l.insertDinos(save.Objects, int(worldId), tx)
	if err != nil {
		return fmt.Errorf("Inserting dino:\n%w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Could not commit SQL transaction:\n%w", err)
	}

	return nil
}

func (l *Loader) insertDinos(objlists [][]*ark.GameObject, world int, tx *sqlx.Tx) error {
	stmt, err := tx.Prepare(`INSERT INTO dinos VALUES (
									?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,
									?,?,?,?,?,?,
									?,?,?,?,?,?,?,?,
									?,?,?,?,?,?,?,?,
									?,?,?,?,?,?,?,?)`)
	if err != nil {
		return fmt.Errorf("Could not prepare SQL insert:\n%w", err)
	}

	for listNum, objlist := range objlists {
		for i, obj := range objlist {
			// TamedOnServerName is a good canary for tamed dinos
			server := obj.Properties.Get("TamedOnServerName", 0)
			if server == nil {
				continue
			}

			name := obj.Properties.GetString("TamedName", 0)
			statsCurrent := make([]float64, 12)
			pointsWild := make([]int64, 12)
			pointsTamed := make([]int64, 12)
			var levelWild int64
			var levelTamed int64

			loc := obj.Location

			var err error
			parentClass := 0
			parentName := ""
			if obj.Parent != nil {
				loc = obj.Parent.Location
				parentClass, err = l.getOrAddClass(tx, obj.Parent.ClassName.Name)
				if err != nil {
					return err
				}

				parentName = obj.Parent.Properties.GetString("BoxName", 0)
				if parentName == "" {
					parentName = obj.Parent.Properties.GetString("PlayerName", 0)
				}
			}

			cscProp := obj.Properties.GetTyped("MyCharacterStatusComponent", 0, ark.ObjectPropertyType)
			if cscProp != nil {
				cscId := cscProp.(*ark.ObjectProperty).Id
				csc := objlist[cscId]

				for index := 0; index < 12; index++ {
					statsCurrent[index] = csc.Properties.GetFloat("CurrentStatusValues", index)
					pointsWild[index] = csc.Properties.GetInt("NumberOfLevelUpPointsApplied", index)
					pointsTamed[index] = csc.Properties.GetInt("NumberOfLevelUpPointsAppliedTamed", index)
				}

				levelWild = csc.Properties.GetInt("BaseCharacterLevel", 0)
				levelTamed = csc.Properties.GetInt("ExtraCharacterLevel", 0)
			}

			dinoId1 := obj.Properties.GetInt("DinoID1", 0)
			dinoId2 := obj.Properties.GetInt("DinoID2", 0)

			colors := make([]int64, 6)
			for i := range colors {
				colors[i] = obj.Properties.GetInt("ColorSetIndices", i)
			}

			classId, err := l.getOrAddClass(tx, obj.ClassName.Name)
			if err != nil {
				return err
			}

			_, err = stmt.Exec(
				i,
				listNum,
				world,
				classId,
				name,
				levelWild,
				levelTamed,
				dinoId1,
				dinoId2,
				obj.IsCryopod,
				parentClass,
				parentName,

				loc.X, loc.Y, loc.Z,

				colors[0], colors[1], colors[2],
				colors[3], colors[4], colors[5],

				statsCurrent[0], statsCurrent[1], statsCurrent[2], statsCurrent[3],
				statsCurrent[4], statsCurrent[7], statsCurrent[8], statsCurrent[9],

				pointsWild[0], pointsWild[1], pointsWild[2], pointsWild[3],
				pointsWild[4], pointsWild[7], pointsWild[8], pointsWild[9],

				pointsTamed[0], pointsTamed[1], pointsTamed[2], pointsTamed[3],
				pointsTamed[4], pointsTamed[7], pointsTamed[8], pointsTamed[9],
			)

			if err != nil {
				return fmt.Errorf("Could not insert object %d:\n%w", i, err)
			}
		}
	}

	return nil
}

func (l *Loader) getOrAddClass(tx *sqlx.Tx, bpName string) (int, error) {
	class := l.classMap.Get(bpName)
	if class == nil {
		class = l.classMap.Add(bpName)
		_, err := tx.Exec("INSERT INTO classes (id, class, name) VALUES (?,?,?)",
			class.Id, bpName, class.Name)
		if err != nil {
			return 0, fmt.Errorf("Adding %s to the class table:\n%w", bpName, err)
		}
	}

	return class.Id, nil
}

func (l *Loader) watcher() {
	for {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatalf("Error creating file watcher:\n%s", err)
		}

		for _, path := range l.saveFiles {
			err = watcher.Add(path)
			if err != nil {
				log.Fatalf("Error watching %s:\n%s", path, err)
			}
		}

		select {
		case event := <-watcher.Events:
			err = watcher.Close()
			if err != nil {
				log.Fatalf("Error closing watcher:\n%s", err)
			}

			time.Sleep(5 * time.Millisecond) // Wait for the rm/rename to finish
			err = l.processSavefile(event.Name)
			if err != nil {
				log.Fatalf("Error reloading save %s:\n%s", err)
			}

		case err := <-watcher.Errors:
			log.Fatalf("Error watching save file:\n%s", err)
		}
	}

}
