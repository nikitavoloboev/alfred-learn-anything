package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// downloadCuratedLists downloads curated lists from GitHub.
func downloadCuratedLists() {
	resp, err := http.Get("https://raw.githubusercontent.com/learn-anything/curated-lists/master/readme.md")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// TODO: add it to cache
	// ioutil.WriteFile("lists.md", body, 0600)
}

// doSearchLists searches curated lists.
func doSearchLists() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	// TODO: where is cache placed?
	parseList("lists.md")

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()

	return nil
}

// parseList parses a markdown list for links.
func parseList(file string) {
	bytes, _ := ioutil.ReadFile(file)

	// Regex to extract markdown links
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	// Read string line by line and apply regex
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		wf.NewItem(matches[0][1]).Arg(matches[0][2]).Valid(true).UID(matches[0][1])
	}
}

func readList(listName string) {

}
