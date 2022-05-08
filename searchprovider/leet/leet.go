package leet

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "io"
    "log"
    "net/http"
    "net/url"
    "strconv"
)

type SearchProvider struct {
}

const baseURL string = "https://www.1377x.to"

// Torrent contains meta information about the torrent
type leetTorrent struct {
	DescURL string
	Name    string
	Size    string
	UplDate string
	// Seeders and Leechers are converted to -1 if cannot be converted to integers
	Seeders  int
	Leechers int
}

func (provider *SearchProvider) Name() string {
	return "1337x"
}

func (provider *SearchProvider) Search(title string, searchpostfixes []string) string {
	for _, searchpostfix := range searchpostfixes {
		torrents, err := search(title + " " + searchpostfix)

		if err != nil {
			log.Printf("1337x error occured: '%s'\n", err.Error())
		}

		if len(torrents) > 0 {
			magnet, err := extractMag(torrents[0].DescURL)
			if err != nil {
				log.Printf("1337x magnet extract error occured: '%s'\n", err.Error())
				continue
			}

			return magnet
		}
	}

	torrents, _ := search(title)

	if len(torrents) > 0 {
		magnet, err := extractMag(torrents[0].DescURL)
		if err != nil {
			log.Printf("1337x magnet extract error occured: '%s'\n", err.Error())
			return ""
		}

		return magnet
	}

	return ""
}

func (provider *SearchProvider) Init() {
	// nothing to be done here
}

// A typical final url looks like:
// https://1337x.to/search/Dumas/1/
func buildSearchURL(in string) (string, error) {
	var URL *url.URL
	URL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("error during url parsing: %v", err)
	}

	URL.Path += "/search/" + in + "/1/"

	return URL.String(), nil
}

func parseSearchPage(htmlReader io.ReadCloser) ([]leetTorrent, error) {
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("could not load html response into GoQuery: %v", err)
	}

	// torrents stores a list of torrents made up of the torrent description url,
	// its name, its size, its upload date, its seeders, and its leechers
	var torrents []leetTorrent

	// Results are located in a clean html <table>
	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		var t leetTorrent

		// Name is the text of the 2nd <a> tag, and desc URL is the href
		path, ok := s.Find("a").Eq(1).First().Attr("href")
		if !ok {
			log.Print("Could not find a description page for a torrent so ignoring it")
			return
		}
		t.DescURL = baseURL + path
		t.Name = s.Find("a").Eq(1).First().Text()

		// Seeders and leechers are located in the 2nd and 3rd <td>.
		// We convert it to integers and if conversion fails we convert it to -1.
		seedersStr := s.Find("td").Eq(1).First().Text()
		seeders, err := strconv.Atoi(seedersStr)
		if err != nil {
			seeders = -1
		}
		t.Seeders = seeders

		leechersStr := s.Find("td").Eq(2).First().Text()
		leechers, err := strconv.Atoi(leechersStr)
		if err != nil {
			leechers = -1
		}
		t.Leechers = leechers

		// Upload date is the text of the 4th <td> tag
		t.UplDate = s.Find("td").Eq(3).First().Text()

		// Size is the text of the 5th <td> tag
		t.Size = s.Find("td").Eq(4).First().Text()

		torrents = append(torrents, t)
	})

	return torrents, nil
}

// Lookup takes a user search as a parameter, launches the http request
// with a custom timeout, and returns clean torrent information fetched from 1337x.to
func search(in string) ([]leetTorrent, error) {
	searchUrl, err := buildSearchURL(in)
	if err != nil {
		return nil, fmt.Errorf("error while building url: %v", err)
	}

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

// parseDescPage parses the torrent description page and extracts the magnet link
func parseDescPage(htmlReader io.ReadCloser) (string, error) {
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return "", fmt.Errorf("could not load html response into GoQuery: %v", err)
	}

	magnet, ok := doc.Find(".torrent-detail-page li a").Eq(0).First().Attr("href")
	if !ok {
		return "", fmt.Errorf("could not extract magnet link")
	}

	return magnet, nil
}

// ExtractMag opens the torrent description page and extracts the magnet link.
// A user timeout is set.
func extractMag(descURL string) (string, error) {
	resp, err := http.Get(descURL)
	if err != nil {
		return "", fmt.Errorf("error while fetching url: %v", err)
	}

	magnet, err := parseDescPage(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while parsing torrent description page: %v", err)
	}

	return magnet, nil
}
