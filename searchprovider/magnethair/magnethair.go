package magnethair

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	searchURL    = "https://magnetdl.hair/api.php?url=/q.php?q=%s"
	magnetFormat = "magnet:?xt=urn:btih:%s&dn=%s&tr=udp://tracker.torrent.eu.org:451/announce&tr=udp://tracker.tiny-vps.com:6969/announce&tr=http://tracker.foreverpirates.co:80/announce&tr=udp://tracker.cyberia.is:6969/announce&tr=udp://exodus.desync.com:6969/announce&tr=udp://explodie.org:6969/announce&tr=udp://tracker.opentrackr.org:1337/announce&tr=udp://9.rarbg.to:2780/announce&tr=udp://tracker.internetwarriors.net:1337/announce&tr=udp://ipv4.tracker.harry.lu:80/announce&tr=udp://open.stealth.si:80/announce&tr=udp://9.rarbg.to:2900/announce&tr=udp://9.rarbg.me:2720/announce&tr=udp://opentor.org:2710/announce"
)

type SearchProvider struct {
}

func (provider *SearchProvider) Name() string {
	return "magnethair"
}

// {"id":"76626463","name":"The Lord of the Rings The Rings of Power S02E06 Where Is He 1080p AMZN WEB-DL DD","info_hash":"D5C51C360011A5EBDAE7A79C22C388947975D1E5","leechers":"3824","seeders":"4311","num_files":"0","size":"2895225325","username":"jajaja","added":"1726730401","status":"vip","category":"208","imdb":"tt7631058"}
type apiResult struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	InfoHash string `json:"info_hash"`
	Seeders  string `json:"seeders"`
	Leechers string `json:"leechers"`
	Size     string `json:"size"`
	Added    string `json:"added"`
}

func (provider *SearchProvider) Search(title string, searchPostfixes []string) string {
	for _, searchPostfix := range searchPostfixes {
		torrents, _ := search(title + " " + searchPostfix)

		if len(torrents) > 0 {
			return fmt.Sprintf(magnetFormat, torrents[0].InfoHash, torrents[0].Name)
		}
	}

	torrents, err := search(title)
	if err != nil {
		log.Printf("error occured during search: %v", err)
	}

	if len(torrents) > 0 {
		return fmt.Sprintf(magnetFormat, torrents[0].InfoHash, torrents[0].Name)
	}

	return ""
}

// Lookup takes a user search as a parameter, launches the http request
// with a custom timeout, and returns clean torrent information fetched from 1337x.to
func search(in string) ([]apiResult, error) {
	searchUrl := buildSearchURL(in)

	resp, err := http.Get(searchUrl)
	if err != nil {
		return nil, fmt.Errorf("error while fetching url: %v", err)
	}

	torrents, err := parseApiRepsonse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing torrent search results: %v", err)
	}

	return torrents, nil
}

func parseApiRepsonse(jsonReader io.ReadCloser) ([]apiResult, error) {
	// parse the json response
	// unmarshal json response
	data, err := io.ReadAll(jsonReader)
	if err != nil {
		return nil, fmt.Errorf("could not read api response: %v", err)
	}
	var results []apiResult
	err = json.Unmarshal(data, &results)
	if err != nil {
		return nil, fmt.Errorf("could not load api response into apiResult struct: %v", err)
	}

	// filter out anything with id < 1
	var torrents []apiResult
	for _, result := range results {
		if result.ID != "0" {
			torrents = append(torrents, result)
		}
	}

	return torrents, nil
}

func buildSearchURL(in string) string {
	// turn in into http url encoded string
	query := url.QueryEscape(in)

	return fmt.Sprintf(searchURL, query)
}

func (provider *SearchProvider) Init() {
	// nothing needs to be done
}
