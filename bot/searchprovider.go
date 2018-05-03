package bot

type SearchProvider interface {
    Search(Title string, SearchPostfixes []string) string
    Name() string
}
