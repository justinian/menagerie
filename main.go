package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

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

	saves, err := findFiles(args)
	if err != nil {
		log.Fatal(err)
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

func findFiles(paths []string) ([]string, error) {
	savePattern := regexp.MustCompile("[A-Z][a-z]+(_P)?.ark$")

	saves := make([]string, 0, len(paths))

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		if !info.IsDir() {
			saves = append(saves, path)
			continue
		}

		err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() && savePattern.MatchString(path) {
				saves = append(saves, path)
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("Searching %s: %w", path, err)
		}
	}

	return saves, nil
}
