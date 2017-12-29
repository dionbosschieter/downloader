package main

import (
    "time"
    "fmt"
    "os/exec"
    "bytes"
    piratebay "github.com/gnur/go-piratebay"
)

var client piratebay.Piratebay
var torrent piratebay.Torrent

func InitClient() {
    client = piratebay.Piratebay { Url: "https://thepiratebay.org" }
}

func Search(title string, location string) {
    Log("Searching for " + title)
    torrents,err := client.Search(title)

    if err != nil {
        Log2Me(err.Error())
        return
    }

    if len(torrents) == 0 {
        return
    }

    AddTorrent(torrents[0].MagnetLink, location)
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

    Log2Me(fmt.Sprintf("Executed subtitle downloader for %s", path))
}

func Log2Me(message string) {
    tbot.Send(&me, message)
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
