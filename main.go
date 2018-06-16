// Command awesome is Alfred 3 workflow for quickly navigating GitHub Awesome lists.
package main

import (
	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// Icons
	iconDefault = &aw.Icon{Value: "icon.png"}
	iconUpdate  = &aw.Icon{Value: "icons/update-available.png"}

	query string

	repo = "nikitavoloboev/alfred-awesome-lists"

	// Workflow stuff
	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	searchAwesomeLists()
}

func main() {
	wf.Run(run)
}
