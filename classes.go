package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var keynames = []string{"items", "species"}
var whitespace = regexp.MustCompile(`\s+`)
var cSuffix = regexp.MustCompile(`_C$`)

type classSpec struct {
	Name      string `json:"name"`
	Blueprint string `json:"bp"`
}

type Class struct {
	Name string
	Id   int
}

type ClassMap struct {
	Map    map[string]Class
	nextId int
}

func (cm *ClassMap) Get(bpName string) *Class {
	if class, ok := cm.Map[bpName]; ok {
		return &class
	}
	return nil
}

func (cm *ClassMap) Add(bpName string) *Class {
	class := Class{
		Name: cSuffix.ReplaceAllString(bpName, ""),
		Id:   cm.nextId,
	}

	cm.Map[bpName] = class
	cm.nextId++

	return &class
}

func readSpecFiles(paths ...string) (*ClassMap, error) {
	classCount := 1 // leave 0 for "none"
	classNames := make(map[string]Class)

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Opening spec file %f:\n%w", path, err)
		}

		jsonData, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("Reading spec file %f:\n%w", path, err)
		}

		var values map[string]json.RawMessage
		if err := json.Unmarshal(jsonData, &values); err != nil {
			return nil, fmt.Errorf("Loading spec file %f:\n%w", path, err)
		}

		for _, key := range keynames {
			if raw, ok := values[key]; ok {
				var specs []classSpec
				if err := json.Unmarshal(raw, &specs); err != nil {
					return nil, fmt.Errorf("Loading specs from file %f:\n%w", path, err)
				}

				for _, spec := range specs {
					parts := strings.Split(spec.Blueprint, ".")
					className := parts[len(parts)-1]
					classNames[className] = Class{
						Name: whitespace.ReplaceAllString(spec.Name, " "),
						Id:   classCount,
					}
					classCount++
				}
			}
		}
	}

	return &ClassMap{Map: classNames, nextId: classCount}, nil
}
