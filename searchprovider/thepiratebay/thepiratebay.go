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
        pbytorrents,_ := provider.client.Search(title + " " + searchpostfix)

        if len(pbytorrents) > 0 {
            return pbytorrents[0].MagnetLink
        }
    }

    pbytorrents,_ := provider.client.Search(title)

    if len(pbytorrents) > 0 {
        return pbytorrents[0].MagnetLink
    }

    return ""
}

func New() (provider *SearchProvider) {
    client := pby.Piratebay{Url: "https://thepiratebay.org"}

    provider = &SearchProvider{
        client: client,
    }

    return
}
