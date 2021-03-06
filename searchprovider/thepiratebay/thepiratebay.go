package thepiratebay

import (
    pby "github.com/gnur/go-piratebay"
)

type SearchProvider struct {
    client pby.Piratebay
}

func (provider *SearchProvider) Name() string {
    return "thepiratebay"
}

func (provider *SearchProvider) Search(title string, searchpostfixes []string) string {
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

func (provider *SearchProvider) Init() {
    provider.client = pby.Piratebay {
        Url: "https://thepiratebay.org",
    }
}