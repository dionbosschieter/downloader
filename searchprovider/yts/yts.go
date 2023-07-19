package yts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	downloadButtonClass = ".magnet-download.download-torrent.magnet"
	searchURL           = "https://yts.pm/ajax/search?query=%s"
)

var searchError = errors.New("search error occured")

type ytsSearchEntry struct {
	Year  string `json:"year"`
	URL   string `json:"url"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

type ytsSearchResult struct {
	Status string           `json:"status"`
	Data   []ytsSearchEntry `json:"data"`
}

type SearchProvider struct {
}

func (provider *SearchProvider) Name() string {
	return "yts"
}

func (provider *SearchProvider) Search(title string, searchPostfixes []string) string {
	for _, searchPostfix := range searchPostfixes {
		movies, _ := search(title + " " + searchPostfix)

		if len(movies) > 0 {
			movie := movies[0]

			return movie.Magnet()
		}
	}

	movies, err := search(title)
	if err != nil {
		log.Printf("error occured during search: %v", err)
	}

	if len(movies) > 0 {
		movie := movies[0]

		return movie.Magnet()
	}

	return ""
}

// Lookup takes a user search as a parameter, launches the http request
// with a custom timeout, and returns clean torrent information fetched from 1337x.to
func search(in string) ([]ytsSearchEntry, error) {
	searchUrl := buildSearchURL(in)

	resp, err := http.Get(searchUrl)
	if err != nil {
		return nil, fmt.Errorf("error while fetching url: %v", err)
	}

	torrents, err := parseSearchPage(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing torrent search results: %v", err)
	}

	return torrents, nil
}

func parseSearchPage(body io.ReadCloser) ([]ytsSearchEntry, error) {
	var results ytsSearchResult
	pageData, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("could not read page data: %v", err)
	}
	err = json.Unmarshal(pageData, &results)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal json: %v", err)
	}

	if results.Status != "ok" {
		return nil, searchError
	}

	return results.Data, nil
}

func (e ytsSearchEntry) Magnet() string {
	searchEntryURL := "https:" + e.URL
	response, err := http.Get(searchEntryURL)
	if err != nil {
		log.Printf("error occured during magnitization of yts url: %s", searchEntryURL)
	}

	magnet, err := parseDescPage(response.Body)

	return magnet
}

// parseDescPage parses the torrent description page and extracts the magnet link
func parseDescPage(htmlReader io.ReadCloser) (string, error) {
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return "", fmt.Errorf("could not load html response into GoQuery: %v", err)
	}

	magnet, ok := doc.Find(downloadButtonClass).Eq(0).First().Attr("href")
	if !ok {
		return "", fmt.Errorf("could not extract magnet link")
	}

	return magnet, nil
}

// A search url looks like:
// https://$baseURL$/ajax/search?query=test
func buildSearchURL(in string) string {
	return fmt.Sprintf(searchURL, url.QueryEscape(in))
}

func (provider *SearchProvider) Init() {
	// nothing needs to be done
}
