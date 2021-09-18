package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/jmoiron/sqlx"
)

type apiHandler struct {
	lock      sync.Mutex
	loader    *Loader
	tamedStmt *sqlx.Stmt
	wildStmt  *sqlx.Stmt
}

func (ah *apiHandler) getDinos(tamed bool) ([]dinoResult, error) {
	ah.loader.lock.Lock()
	defer ah.loader.lock.Unlock()

	db := ah.loader.db

	query := getWildDinos
	stmt := ah.wildStmt
	if tamed {
		query = getTamedDinos
		stmt = ah.tamedStmt
	}

	var err error
	if stmt == nil {
		stmt, err = db.Preparex(query)
		if err != nil {
			return nil, fmt.Errorf("Database Error: %w", err)
		}

		if tamed {
			ah.tamedStmt = stmt
		} else {
			ah.wildStmt = stmt
		}
	}

	var result []dinoResult
	err = stmt.Select(&result)
	if err != nil {
		return nil, fmt.Errorf("Database Error: %w", err)
	}

	return result, nil
}

func (ah *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing query: %s", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tamed := r.Form.Get("type") != "wild"
	result, err := ah.getDinos(tamed)
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
	apiHandler := &apiHandler{loader: loader}

	sm := http.NewServeMux()
	sm.Handle("/api/dinos", apiHandler)

	fs := http.FileServer(http.Dir("static"))
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/tamed.html", http.StatusFound)
			return
		}
		fs.ServeHTTP(w, r)
	})

	log.Printf("Listening on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, loggingWrapper(sm)))
}
