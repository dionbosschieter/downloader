package yts

import (
    ytsclient "github.com/qopher/ytsgo"
)

type SearchProvider struct {
    client ytsclient.Client
}

func (provider *SearchProvider) Name() string {
    return "yts"
}

func (provider *SearchProvider) Search(title string, searchPostfixes []string) string {
    for _, searchPostfix := range searchPostfixes {
        movies,_ := provider.client.ListMovies(ytsclient.LMSearch(title + " " + searchPostfix))

        if len(movies.Movies) > 0 {
            movie := movies.Movies[0]

            if len(movie.Torrents) > 0 {
                return movie.Torrents[0].Magnet()
            }
        }
    }

    movies,_ := provider.client.ListMovies(ytsclient.LMSearch(title))

    if len(movies.Movies) > 0 {
        movie := movies.Movies[0]

        if len(movie.Torrents) > 0 {
            return movie.Torrents[0].Magnet()
        }
    }

    return ""
}

func (provider *SearchProvider) Init() {
    client, _ := ytsclient.New()

    provider.client = *client
}