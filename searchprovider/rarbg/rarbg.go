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

func (provider *SearchProvider) Search(title string, searchPostfixes []string) (magnet string) {
    // query with the provided postfixes
    for _, searchPostfix := range searchPostfixes {
        var query = map[string]string{"search_string":title + " " + searchPostfix, "sort": "seeders"}

        result,_ := provider.client.Search(query)

        if len(result.Torrents) > 0 {
            return result.Torrents[0].MagnetURL
        }
    }

    // fall back to postfix less search
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
