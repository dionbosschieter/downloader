package main

import (
    "github.com/tubbebubbe/transmission"
    "fmt"
    "time"
)

var tclient transmission.TransmissionClient

func GetTorrents() string {
    returnstring := "Current torrents:\n"

    torrents, err := tclient.GetTorrents()
    if err != nil {
        panic(err)
    }

    for _, t := range torrents {
         returnstring += t.Name + " " + fmt.Sprintf("%.4f", t.PercentDone * 100) + "%\n"
    }

    return returnstring
}

func AddTorrent(magnet string, location string) {
    cmd,_ := transmission.NewAddCmdByMagnet(magnet)
    cmd.SetDownloadDir(location)
    add,err := tclient.ExecuteAddCommand(cmd)

    if err != nil {
        Log2Me(err.Error())
        return
    }

    Log2Me("added torrent: " + add.Name)
    go WaitTillFinished(add)
}

func WaitTillFinished(add transmission.TorrentAdded) {
    for {
        if TorrentIsFinished(add.ID) {
            Log2Me(add.Name + " is finished")
            break
        }

        time.Sleep(time.Second * 5)
    }

    path := GetTorrentPath(add.ID)
    if path != "" {
        DownloadSubtitles(path)
    }
    RemoveTorrent(add.ID)
}

func GetTorrentPath(id int) string {
    torrents, err := tclient.GetTorrents()
    if err != nil {
        Log2Me(err.Error())
        return ""
    }

    for _,torrent := range torrents {
        if torrent.ID == id {
            return torrent.DownloadDir
        }
    }

    Log2Me(fmt.Sprintf("Cant find torrent by id %d", id))
    return ""

}

// remove but keep files
func RemoveTorrent(id int) {
    cmd,_ := transmission.NewDelCmd(id, false)

    tclient.ExecuteCommand(cmd)
}

func TorrentIsFinished(id int) bool {
    torrents, err := tclient.GetTorrents()
    if err != nil {
        Log2Me(err.Error())
        return false
    }

    for _,torrent := range torrents {
        if torrent.ID == id {
            return torrent.PercentDone == float64(1)
        }
    }

    Log2Me(fmt.Sprintf("Cant find torrent by id %d", id))
    return true
}

func SetupTransmissionClient() {
    tclient = transmission.New("http://<host>:<port>", "", "")
}
