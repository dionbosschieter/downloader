package main

import (
    pby "github.com/gnur/go-piratebay"
)

type searchprovider struct {
    client pby.Piratebay
}

var SearchProvider searchprovider

func (provider *searchprovider) Name() string {
    return "thepiratebay"
}

func (provider *searchprovider) Search(title string, searchpostfixes []string) string {
    for _, searchpostfix := range searchpostfixes {
        torrents,_ := provider.client.Search(title + " " + searchpostfix)

        if len(torrents) > 0 {
            return torrents[0].MagnetLink
        }
    }

    torrents,_ := provider.client.Search(title)

    if len(torrents) > 0 {
        return torrents[0].MagnetLink
    }

    return ""
}

func (provider *searchprovider) Init() {
    provider.client = pby.Piratebay {
        Url: "https://thepiratebay.org",
    }
}