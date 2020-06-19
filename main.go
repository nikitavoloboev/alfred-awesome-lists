// Command awesome is Alfred 3 workflow for quickly navigating GitHub Awesome lists.
package main

import (
	"log"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// Icons
	iconUpdate = &aw.Icon{Value: "icons/update-available.png"}

	query string

	repo = "nikitavoloboev/alfred-awesome-lists"

	// Workflow stuff
	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	reload := func() (interface{}, error) {
		listIndex, err := searchAwesomeLists()

		if err != nil {
			log.Fatal(err)
		}

		var lists []List
		for _, list := range listIndex {
			lists = append(lists, list)
		}

		return lists, err
	}

	// Cache Awesome lists for 6 hours
	maxCache := 6 * time.Hour
	var lists []List
	err := wf.Cache.LoadOrStoreJSON("awesomeLists", maxCache, reload, &lists)

	if err != nil {
		wf.Fatal(err.Error())
	}

	for _, list := range lists {
		wf.NewItem(list.Name).UID(list.UID).Valid(true).Arg(list.URL)
	}

	args := wf.Args()
	var searchQuery string
	if len(args) > 0 {
		searchQuery = args[0]
	}

	if searchQuery == "" {
		wf.WarnEmpty("No matching items", "Try a different query?")
	} else {
		wf.Filter(searchQuery)
		wf.WarnEmpty("No matching items", "Try a different query?")
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
