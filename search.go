package main

import (
	"log"

	parser "github.com/nikitavoloboev/markdown-parser"
)

// List holds an awesome list.
type List struct {
	UID  string
	Name string
	URL  string
}

func searchAwesomeLists() (map[string]List, error) {
	showUpdateStatus()

	log.Printf("query=%s", query)

	// Get the list from GitHub
	urls, err := parser.ParseMarkdownURL("https://raw.githubusercontent.com/sindresorhus/awesome/main/readme.md")
	if err != nil {
		log.Println("Error parsing links")
	}

	links := make(map[string]List)

	for name, url := range urls {
		links[url] = List{
			UID:  name,
			Name: name,
			URL:  url,
		}
	}

	return links, nil
}
