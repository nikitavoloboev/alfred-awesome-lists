package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/docopt/docopt-go"
	"github.com/nikitavoloboev/markdown-parser/parser"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var usage = `alfred-awesome-lists [search|check] [<query>]

Access any awesome list from Alfred in seconds.

Usage:
	alfred-awesome-lists search [<query>]
	alfred-awesome-lists check
    alfred-awesome-lists -h

Options:
    -h, --help    Show this message and exit.
`

var (
	iconAvailable = &aw.Icon{Value: "icons/update.png"}
	cacheName     = "awesome-lists.json"
	maxCacheAge   = 180 * time.Minute

	repo = "nikitavoloboev/alfred-awesome-lists"
	wf   *aw.Workflow
)

func init() {
	// TODO: add update.GitHub(repo) later
	wf = aw.New(update.GitHub(repo))
}

func run() {
	// Pass wf.Args() to docopt because our update logic relies on
	// AwGo's magic actions.
	args, _ := docopt.Parse(usage, wf.Args(), true, wf.Version(), false, true)

	// alternate action: get available releases from remote
	if args["check"] != false {
		wf.TextErrors = true
		log.Println("checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		return
	}

	// _script filter
	var query string
	if args["<query>"] != nil {
		query = args["<query>"].(string)
	}

	log.Printf("query=%s", query)

	// call self with "check" command if an update is due and a
	// check job isn't already running.
	if wf.UpdateCheckDue() && !aw.IsRunning(updateJobName) {
		log.Println("running update check in background...")
		cmd := exec.Command("./alfred-awesome-lists", "check")
		if err := aw.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("error starting update check: %s", err)
		}
	}

	if query == "" { // Only show update status if query is empty
		// Send update status to Alfred
		if wf.UpdateAvailable() {
			wf.NewItem("update available!").
				Subtitle("↩ to install").
				Autocomplete("workflow:update").
				Valid(false).
				Icon(iconAvailable)
		}
	}

	// TODO: add cache
	// if wf.Cache.Exists(cacheName) {
	// 	if err := wf.Cache.LoadJSON(cacheName, )
	// }

	// parse URL for links
	links, err := parser.ParseMarkdownURL("https://raw.githubusercontent.com/sindresorhus/awesome/master/readme.md")
	if err != nil {
		log.Println("Error parsing links")
	}

	// add all links to Alfred
	for k, v := range links {
		wf.NewItem(k).Arg(v).Valid(true)
	}

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
