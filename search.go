package main

import (
	"log"
	"os"

	"github.com/nikitavoloboev/markdown-parser"
)

func searchAwesomeLists() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	// Get the list from GitHub
	links, err := parser.ParseMarkdownURL("https://raw.githubusercontent.com/sindresorhus/awesome/master/readme.md")
	if err != nil {
		log.Println("Error parsing links")
	}

	// Add all links to Alfred
	for k, v := range links {
		wf.NewItem(k).Arg(v + "#readme").Valid(true).UID(k)
	}

	// TODO: add cache

	// if err := wf.Session.LoadOrStoreJSON("awesome", getWins, &wins); err != nil {
	// 	return nil, err
	// }

	query = os.Args[1]

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()

	return nil
}
