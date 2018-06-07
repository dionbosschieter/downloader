package main

import (
    rar "github.com/ricksancho/rarbg-torrentapi"
)

type searchprovider struct {
    client *rar.Client
}

var SearchProvider searchprovider

func (provider *searchprovider) Name() string {
    return "rarbg"
}

func (provider *searchprovider) Search(title string, searchpostfixes []string) (magnet string) {
    var query = map[string]string{"search_string":title, "sort": "seeders"}

    result,_ := provider.client.Search(query)

    if len(result.Torrents) > 0 {
        return result.Torrents[0].MagnetURL
    }

    return ""
}

func (provider *searchprovider) Init() {
    provider.client,_ = rar.New(1337)
    provider.client.Init()
}
