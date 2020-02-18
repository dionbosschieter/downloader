package main

import (
    yts "github.com/qopher/ytsgo"
)

type searchprovider struct {
    client yts.Client
}

var SearchProvider searchprovider

func (provider *searchprovider) Name() string {
    return "yts"
}

func (provider *searchprovider) Search(title string, searchPostfixes []string) string {
    for _, searchPostfix := range searchPostfixes {
        movies,_ := provider.client.ListMovies(yts.LMSearch(title + " " + searchPostfix))

        if len(movies.Movies) > 0 {
            movie := movies.Movies[0]

            if len(movie.Torrents) > 0 {
                return movie.Torrents[0].Magnet()
            }
        }
    }

    movies,_ := provider.client.ListMovies(yts.LMSearch(title))

    if len(movies.Movies) > 0 {
        movie := movies.Movies[0]

        if len(movie.Torrents) > 0 {
            return movie.Torrents[0].Magnet()
        }
    }

    return ""
}

func (provider *searchprovider) Init() {
    client, _ := yts.New()

    provider.client = *client
}