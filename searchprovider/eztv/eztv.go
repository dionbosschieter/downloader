package eztv

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	searchURL       = "https://eztv1.xyz/search/%s"
	showMagnetLinks = "layout=def_wlinks"
)

type searchEntry struct {
	URL     string
	Magnet  string
	Title   string
	Size    string
	UplDate string
	Seeders int
}

type SearchProvider struct {
}

func (provider *SearchProvider) Name() string {
	return "eztv"
}

func (provider *SearchProvider) Search(title string, searchPostfixes []string) string {
	for _, searchPostfix := range searchPostfixes {
		torrents, _ := search(title + " " + searchPostfix)

		if len(torrents) > 0 {
			return torrents[0].Magnet
		}
	}

	torrents, err := search(title)
	if err != nil {
		log.Printf("error occured during search: %v", err)
	}

	if len(torrents) > 0 {
		return torrents[0].Magnet
	}

	return ""
}

// Lookup takes a user search as a parameter, launches the http request
// with a custom timeout, and returns clean torrent information fetched from eztv
func search(in string) ([]searchEntry, error) {
	searchUrl := buildSearchURL(in)

	resp, err := http.Post(searchUrl, "application/x-www-form-urlencoded", strings.NewReader(showMagnetLinks))
	if err != nil {
		return nil, fmt.Errorf("error while fetching url: %v", err)
	}

	torrents, err := parseSearchPage(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing torrent search results: %v", err)
	}

	return torrents, nil
}

func parseSearchPage(htmlReader io.ReadCloser) ([]searchEntry, error) {
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("could not load html response into GoQuery: %v", err)
	}

	// torrents stores a list of torrents made up of the torrent description url,
	// its name, its size, its upload date, its seeders, and its leechers
	var torrents []searchEntry

	// Results are located in a clean html <table>
	doc.Find("table.forum_header_border > tbody > tr.forum_header_border").Each(func(i int, s *goquery.Selection) {
		var t searchEntry

		// Magnet is the href of the 1st <a> tag in the td.m
		magnetURL, ok := s.Find("td.forum_thread_post > a").Eq(2).First().Attr("href")
		if !ok {
			log.Print("Could not find a magnet URL for a torrent so ignoring it")
			return
		}
		t.Magnet = magnetURL
		t.Title, _ = s.Find("td.forum_thread_post > a").Eq(1).First().Attr("title")

		// Seeders and leechers are located in the 2nd and 3rd <td>.
		// We convert it to integers and if conversion fails we convert it to -1.
		seedersStr := s.Find("td.forum_thread_post_end > font").Eq(0).First().Text()
		seeders, err := strconv.Atoi(seedersStr)
		if err != nil {
			seeders = -1
		}
		t.Seeders = seeders

		// Upload date is the text of the 3th <td> tag
		t.UplDate = s.Find("td.forum_thread_post").Eq(4).First().Text()

		// Size is the text of the 6th <td> tag
		t.Size = s.Find("td.forum_thread_post").Eq(3).First().Text()

		torrents = append(torrents, t)
	})

	return torrents, nil
}

func buildSearchURL(in string) string {
	query := strings.ReplaceAll(strings.ToLower(in), " ", "-")

	return fmt.Sprintf(searchURL, query)
}

func (provider *SearchProvider) Init() {
	// nothing needs to be done
}
