package bot

import (
    "github.com/dionbosschieter/downloader/searchprovider"
    "gopkg.in/tucnak/telebot.v2"
)

type Query struct {
	Title     string
	Requester *telebot.User
	Path      string
	Magnet    string
	IsMovie   bool
}

func (query *Query) Perform(providers []searchprovider.SearchProvider, postfixes []string) {
    for _,provider := range providers {
        Log("Searching for " + query.Title + " with provider " + provider.Name())
        query.Magnet = provider.Search(query.Title, postfixes)

        if query.Magnet != "" {
            Log("Downloading magnet: " + query.Magnet)
            query.Download()
            break
        }
    }

    if query.Magnet == "" {
        Log2Sender(query.Requester, "Could not find any result for " + query.Title)
    }
}