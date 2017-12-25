package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
)

// doSearch searches all Learn Anything topics.
func doSearch() error {
	log.Printf("query=%s", query)

	m, err := loadValues("maps.json")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range m {
		// log.Printf(strconv.Itoa(k))
		// log.Printf(v)
		wf.NewItem(v).Arg(strconv.Itoa(k)).Valid(true).UID(v)
	}
	return nil
}

type Result struct {
	ID  int    `json:"mapID"`
	Key string `json:"key"`
}

// loadVaules returns ID's and keys from read JSON file.
func loadValues(fileName string) (map[int]string, error) {
	file, err := os.Open(fileName)
	m := make(map[int]string)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(file)
	for {
		var ret Result
		err := dec.Decode(&ret)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
			return m, err
		}
		m[ret.ID] = ret.Key
	}
	return m, nil
}
