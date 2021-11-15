package magnet

import "strings"

type SearchProvider struct {
    // SearchProviders usually have a client
}

func (provider *SearchProvider) Name() string {
    return "magnet"
}

func (provider *SearchProvider) Search(title string, searchPostfixes []string) string {
    if strings.Index(title, "magnet:") == 0 {
        return title
    }

    return ""
}

func (provider *SearchProvider) Init() {
    // we don't need any initializing
}