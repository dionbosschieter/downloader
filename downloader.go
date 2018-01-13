package main

import (
    "time"
    "fmt"
    "os/exec"
    "bytes"
    piratebay "github.com/gnur/go-piratebay"
    "gopkg.in/tucnak/telebot.v2"
)

var client piratebay.Piratebay
var torrent piratebay.Torrent

const (
    telegramToken = "<your-telegram-token>"
    transmissionUrl = "http://<host>:<port>"
    seriePath = "/path"
    moviePath = "/path"
    masterChatId = 1337
)


type DownloadQuery struct {
    Title string
    Requester *telebot.User
    Path string
    Magnet string
}

func InitClient() {
    client = piratebay.Piratebay { Url: "https://thepiratebay.org" }
}

func (q *DownloadQuery) Perform() {
    Log("Searching for " + q.Title)
    torrents,err := client.Search(q.Title)

    if err != nil {
        Log2Me(err.Error())
        return
    }

    if len(torrents) == 0 {
        return
    }

    q.Magnet = torrents[0].MagnetLink
    q.Download()
}

func DownloadSubtitles(path string) {
    cmd := exec.Command("/usr/bin/ruby", "subdbdownloader.rb", path)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()

    if err != nil {
        Log2Me(err.Error())
        return
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
    fmt.Println(date.Format(time.RFC3339),message)
}

func main() {
    InitClient()
    SetupTransmissionClient()
    Log("Init piratebay and transmission client")
    SetupTalkyBot()
}
