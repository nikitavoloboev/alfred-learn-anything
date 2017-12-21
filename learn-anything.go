package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Suggestion holds a suggested map from Learn Anything.
type Suggestion struct {
	Name string `json:"key"`
	ID   string `json:"id"`
}

func doSearch() error {
	id, name, err := getRandomSuggestion()
	if err != nil {
		log.Fatal(err)
	}
	wf.NewItem(name).Arg(id)
	wf.SendFeedback()
	return nil
}

// getRandomSuggestion returns random suggestion from Learn Anything.
func getRandomSuggestion() (id string, name string, err error) {
	res, err := http.Get("https://learn-anything.xyz/api/maps/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var suggestion []Suggestion
	err = json.Unmarshal(body, &suggestion)
	if err != nil {
		log.Fatal(err)
	}
	return suggestion[0].ID, suggestion[0].Name, nil
}
