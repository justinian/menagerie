package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type api struct {
	lock       sync.Mutex
	loader     *Loader
	tamedStmt  *sqlx.Stmt
	wildStmt   *sqlx.Stmt
	worldsStmt *sqlx.Stmt
}

func (ah *api) getTames(w http.ResponseWriter, r *http.Request) {
	ah.loader.lock.Lock()
	defer ah.loader.lock.Unlock()

	db := ah.loader.db

	var err error
	if ah.tamedStmt == nil {
		ah.tamedStmt, err = db.Preparex(getTamedDinos)
		if err != nil {
			log.Printf("Error preparing statement: %s", err)
			http.Error(w, "Database Error", http.StatusInternalServerError)
			return
		}
	}

	var result []dinoResult
	err = ah.tamedStmt.Select(&result)
	if err != nil {
		log.Printf("Error loading dinos: %s", err)
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshalling result: %s", err)
		http.Error(w, "JSON Error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (ah *api) getWorlds(w http.ResponseWriter, r *http.Request) {
	ah.loader.lock.Lock()
	defer ah.loader.lock.Unlock()

	db := ah.loader.db

	var err error
	if ah.worldsStmt == nil {
		ah.worldsStmt, err = db.Preparex(getWorlds)
		if err != nil {
			log.Printf("Error preparing statement: %s", err)
			http.Error(w, "Database Error", http.StatusInternalServerError)
			return
		}
	}

	var result []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	err = ah.worldsStmt.Select(&result)
	if err != nil {
		log.Printf("Error loading dinos: %s", err)
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshalling result: %s", err)
		http.Error(w, "JSON Error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (ah *api) getWild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	worldId := vars["world"]

	ah.loader.lock.Lock()
	defer ah.loader.lock.Unlock()

	db := ah.loader.db

	var err error
	if ah.wildStmt == nil {
		ah.wildStmt, err = db.Preparex(getWildDinos)
		if err != nil {
			log.Printf("Error preparing statement: %s", err)
			http.Error(w, "Database Error", http.StatusInternalServerError)
			return
		}
	}

	var result []dinoResult
	err = ah.wildStmt.Select(&result, worldId)
	if err != nil {
		log.Printf("Error loading dinos: %s", err)
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshalling result: %s", err)
		http.Error(w, "JSON Error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
func runServer(loader *Loader, addr string) {
	api := &api{loader: loader}

	r := mux.NewRouter()
	r.HandleFunc("/api/tames", api.getTames)
	r.HandleFunc("/api/worlds", api.getWorlds)
	r.HandleFunc("/api/worlds/{world}/dinos", api.getWild)

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/tamed.html", http.StatusFound)
	})

	log.Printf("Listening on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, loggingWrapper(r)))
}
