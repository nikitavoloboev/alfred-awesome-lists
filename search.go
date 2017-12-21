package main

import (
	"log"

	"github.com/nikitavoloboev/markdown-parser/parser"
)

func doSearch() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	// Parse URL for links
	links, err := parser.ParseMarkdownURL("https://raw.githubusercontent.com/sindresorhus/awesome/master/readme.md")
	if err != nil {
		log.Println("Error parsing links")
	}

	// Add all links to Alfred
	for k, v := range links {
		wf.NewItem(k).Arg(v).Valid(true).UID(k)
	}

	// TODO: Add cache

	// if err := wf.Session.LoadOrStoreJSON("awesome", getWins, &wins); err != nil {
	// 	return nil, err
	// }

	if query != "" {
		res := wf.Filter(query)
		log.Printf("%d results for `%s`", len(res), query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()

	return nil
}
