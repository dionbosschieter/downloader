package main

import (
	"fmt"
	piratebay "github.com/gnur/go-piratebay"
	rarbg "github.com/ricksancho/rarbg-torrentapi"
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

var client piratebay.Piratebay
var torrent piratebay.Torrent
var rar *rarbg.Client
var settings Settings

type DownloadQuery struct {
	Title     string
	Requester *telebot.User
	Path      string
	Magnet    string
	IsMovie   bool
}

func InitClient() {
	client = piratebay.Piratebay{Url: "https://thepiratebay.org"}
	rar, _ = rarbg.New(1337)
	rar.Init()
}

// todo: define providers in yaml
// todo: loop over providers
// todo: prioritize 720 1080p
func (q *DownloadQuery) Perform() {
	Log("Searching for " + q.Title)

	var rarbgQuery = map[string]string{"search_string": q.Title, "sort": "seeders", "category": "tv"}
	if q.IsMovie {
		rarbgQuery["category"] = "movies"
	}
	pbytorrents, pbyerr := client.Search(q.Title)
	rarresult, rarerr := rar.Search(rarbgQuery)

	if pbyerr != nil {
		Log2Me(pbyerr.Error())
	}
	if rarerr != nil {
		Log2Me(rarerr.Error())
		return
	}

	if len(pbytorrents) == 0 && len(rarresult.Torrents) == 0 {
		Log2Sender(q.Requester, "Could not find any result for: "+q.Title)
		return
	}

	if len(rarresult.Torrents) > 0 {
		q.Magnet = rarresult.Torrents[0].MagnetURL
		q.Download()
	} else {
		q.Magnet = pbytorrents[0].MagnetLink
		q.Download()
	}
}

func Log2Me(message string) {
	tbot.Send(&me, message)
	Log(message)
}

func Log2Sender(sender *telebot.User, message string) {
	tbot.Send(sender, message)
	Log(message)
}

func Log(message string) {
	date := time.Now()
	fmt.Println(date.Format(time.RFC3339), message)
}

func main() {
	if settings.FileExists() {
		settings.Parse()
	} else {
		panic("No settings.yaml is defined, see example.yaml")
	}

	InitClient()
	SetupTransmissionClient()
	Log("Init downloader")
	SetupTalkyBot()
}
