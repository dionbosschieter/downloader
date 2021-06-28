package rarbg

import (
    rar "github.com/ricksancho/rarbg-torrentapi"
)

type SearchProvider struct {
    client *rar.Client
}

func (provider *SearchProvider) Name() string {
    return "rarbg"
}

func (provider *SearchProvider) Search(title string, searchpostfixes []string) (magnet string) {
    var query = map[string]string{"search_string":title, "sort": "seeders"}

    result,_ := provider.client.Search(query)

    if len(result.Torrents) > 0 {
        return result.Torrents[0].MagnetURL
    }

    return ""
}

func (provider *SearchProvider) Init() {
    provider.client,_ = rar.New(1337)
    provider.client.Init()
}
