package bitmagnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	searchURL = "http://localhost:3333/graphql"
	queryFmt  = `query {
  torrentContent {
    search(
      input: {
        queryString: "%s"
        limit: 10
        page: 1
        hasNextPage: true
      }
    ) {
      totalCount
      hasNextPage
      items {
        title
        contentType
        seeders
        leechers
        episodes {
          label
          seasons {
            season
            episodes
          }
        }
        torrent {
          magnetUri
        }
        content {
          adult
          originalTitle
          releaseYear
        }
      }
    }
  }
}`
	minSeeders = 2
)

type SearchProvider struct {
}

func (provider *SearchProvider) Name() string {
	return "bitmagnet"
}

// { "title": "Citadel Honey Bunny (2024) S01E01", "contentType": "tv_show", "seeders": 2, "leechers": 0,
//
//	"episodes": { "label": "S01E01", "seasons": [ { "season": 1, "episodes": [ 1 ] } ] }, "torrent": {
//	  "magnetUri": "magnet:?xt=urn:btih:xxxx&dn=xx.xxx.Bunny.xxxx" }, "content": { "adult": null, "originalTitle": "Citadel Honey Bunny", "releaseYear": 2024 } },
type Torrent struct {
	Title       string `json:"title"`
	ContentType string `json:"contentType"`
	Seeders     int    `json:"seeders"`
	Leechers    int    `json:"leechers"`
	Torrent     struct {
		MagnetUri string `json:"magnetUri"`
	} `json:"torrent"`
	Content struct {
		OriginalTitle string `json:"originalTitle"`
		ReleaseYear   int    `json:"releaseYear"`
		Adult         *bool  `json:"adult"`
	} `json:"content"`
}

type apiResult struct {
	Data struct {
		TorrentContent struct {
			Search struct {
				TotalCount  int       `json:"totalCount"`
				HasNextPage bool      `json:"hasNextPage"`
				Items       []Torrent `json:"items"`
			} `json:"search"`
		} `json:"torrentContent"`
	} `json:"data"`
}

type searchQuery struct {
	Query string `json:"query"`
}

func (provider *SearchProvider) Search(title string, searchPostfixes []string) string {
	for _, searchPostfix := range searchPostfixes {
		torrents, _ := search(title + " " + searchPostfix)

		if len(torrents) > 0 {
			return torrents[0].Torrent.MagnetUri
		}
	}

	torrents, err := search(title)
	if err != nil {
		log.Printf("error occured during search: %v", err)
	}

	if len(torrents) > 0 {
		return torrents[0].Torrent.MagnetUri
	}

	return ""
}

// Lookup takes a user search as a parameter, launches the http request
// with a custom timeout, and returns torrent information fetched from bitmagnet
func search(in string) ([]Torrent, error) {
	graphqlQuery := buildSearchQuery(in)

	searchQueryInput, err := json.Marshal(graphqlQuery)

	resp, err := http.Post(searchURL, "application/json", bytes.NewReader(searchQueryInput))
	if err != nil {
		return nil, fmt.Errorf("error while fetching url: %v", err)
	}

	torrents, err := parseApiRepsonse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing torrent search results: %v", err)
	}

	return torrents, nil
}

func parseApiRepsonse(jsonReader io.ReadCloser) ([]Torrent, error) {
	// parse the json response
	data, err := io.ReadAll(jsonReader)
	if err != nil {
		return nil, fmt.Errorf("could not read api response: %v", err)
	}
	var result apiResult
	// unmarshal json response
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("could not load api response into Torrent struct: %v", err)
	}

	// filter out anything with id < 1
	var torrents []Torrent
	for _, result := range result.Data.TorrentContent.Search.Items {
		if result.Seeders >= minSeeders && (result.ContentType == "tv_show" || result.ContentType == "movie") {
			torrents = append(torrents, result)
		}
	}

	return torrents, nil
}

func buildSearchQuery(title string) searchQuery {
	return searchQuery{
		Query: fmt.Sprintf(queryFmt, title),
	}
}

func (provider *SearchProvider) Init() {
	// nothing needs to be done
}
