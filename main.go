package main

import (
	"fmt"

	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Defaults for Kingpin flags
const (
	defaultMaxResults = "100"
)

// Icons
var (
	IconDefault = &aw.Icon{Value: "icon.png"}
	IconUpdate  = &aw.Icon{Value: "icons/update-available.png"}
)

var (
	// Kingpin and script options
	app *kingpin.Application

	// Application commands
	searchTopicsCmd *kingpin.CmdClause
	searchListsCmd  *kingpin.CmdClause
	updateCmd       *kingpin.CmdClause
	testCmd         *kingpin.CmdClause

	// Script options (populated by Kingpin application)
	query string

	repo = "nikitavoloboev/alfred-learn-anything"

	// Workflow stuff
	wf *aw.Workflow
)

// Sets up kingpin commands
func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
	app = kingpin.New("learn anything", "Search Learn Anything.")

	updateCmd = app.Command("update", "Check for new version.")

	searchTopicsCmd = app.Command("topics", "Search Learn Anything topics.")

	searchListsCmd = app.Command("lists", "Search curated lists.")

	for _, cmd := range []*kingpin.CmdClause{
		searchTopicsCmd, searchListsCmd,
	} {
		cmd.Flag("query", "Search query.").Short('q').StringVar(&query)
	}
}

func run() {
	var err error

	cmd, err := app.Parse(wf.Args())
	if err != nil {
		wf.FatalError(err)
	}

	switch cmd {
	case searchTopicsCmd.FullCommand():
		err = doSearchTopics()
	case searchListsCmd.FullCommand():
		err = doSearchLists()
	case updateCmd.FullCommand():
		err = doUpdate()
	default:
		err = fmt.Errorf("Uknown command: %s", cmd)
	}

	// Check for update
	if err == nil && cmd != updateCmd.FullCommand() {
		err = checkForUpdate()
	}

	if err != nil {
		wf.FatalError(err)
	}
}

func doTest() error {
	return nil
}

func main() {
	wf.Run(run)
}
