package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	var output string
	var address string
	var specfiles []string
	pflag.StringVarP(&output, "out", "o", "ark.db", "Filename of the database to create")
	pflag.StringVarP(&address, "addr", "a", "[::]:8090", "Address to listen on")
	pflag.StringArrayVarP(&specfiles, "spec", "s", nil, "JSON species/item files to load")
	pflag.Parse()

	args := pflag.Args()

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <savefile> ...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "       %s -h  for help\n", os.Args[0])
		os.Exit(1)
	}

	saves := make([]string, 0, len(args))
	for _, savepath := range args {
		info, err := os.Stat(savepath)
		if err != nil {
			log.Fatalf("%s: %s", savepath, err)
		}

		if !info.IsDir() {
			saves = append(saves, savepath)
			continue
		}

		entries, err := os.ReadDir(savepath)
		if err != nil {
			log.Fatalf("Directory %s: %s", savepath, err)
		}

		for _, ent := range entries {
			if ent.IsDir() {
				continue
			}
			if strings.HasSuffix(ent.Name(), ".ark") {
				saves = append(saves, path.Join(savepath, ent.Name()))
			}
		}
	}

	log.Print("Menagerie starting.")
	for _, save := range saves {
		log.Printf("Using save file: %s", save)
	}

	loader, err := createLoader(output, specfiles, saves)
	if err != nil {
		log.Fatal(err)
	}
	defer loader.db.Close()

	err = loader.run()
	if err != nil {
		log.Fatal(err)
	}

	runServer(loader, address)
}
