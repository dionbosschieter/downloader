package bot

import (
	"github.com/dionbosschieter/downloader/searchprovider"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

type Query struct {
	Title     string
	Requester *telebot.User
	Path      string
	Magnet    string
	IsMovie   bool
}

func (query *Query) Perform(providers []searchprovider.SearchProvider, postfixes []string) {
	query.QueryProviders(providers, postfixes)

	if query.Magnet == "" {
		query.Log2Requester("Could not find any result for " + query.Title)
	}
}

func (query *Query) QueryProviders(providers []searchprovider.SearchProvider, postfixes []string) {
	for _, provider := range providers {
		log.Println("Searching for " + query.Title + " with provider " + provider.Name())
		query.Magnet = provider.Search(query.Title, postfixes)

		if query.Magnet != "" {
			log.Println("Downloading magnet: " + query.Magnet)
			query.Download()
			break
		}
	}
}

func (query *Query) Log2Requester(message string) {
	tbot.Send(query.Requester, message)
	log.Println(message)
}

func (query *Query) Schedule(providers []searchprovider.SearchProvider, postfixes []string) {
	for {
		log.Printf("Searching for %s every hour", query.Title)
		query.QueryProviders(providers, postfixes)

		if query.Magnet != "" {
			break
		}

		time.Sleep(1 * time.Hour)
	}
}
