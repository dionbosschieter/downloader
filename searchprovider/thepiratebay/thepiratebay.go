package thepiratebay

type SearchProvider struct {}

func (SearchProvider) Name() string {
    return "thepiratebay"
}

// returns Magnet of first match on rarbg
func (SearchProvider) Search(title string, searchpostfixes []string) string {
    return ""
}
