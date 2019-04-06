package bot

import (
	"fmt"
	"github.com/tubbebubbe/transmission"
	"time"
)

var tclient transmission.TransmissionClient

func GetTorrents() string {
	returnstring := "Current torrents:\n"

	torrents, err := tclient.GetTorrents()
	if err != nil {
		return "Error getting current torrents: " + err.Error()
	}

	for _, t := range torrents {
		returnstring += t.Name + " " + fmt.Sprintf("%.4f", t.PercentDone*100) + "%\n"
	}

	return returnstring
}

func (q *DownloadQuery) Download() {
	cmd, _ := transmission.NewAddCmdByMagnet(q.Magnet)
	cmd.SetDownloadDir(q.Path)
	add, err := tclient.ExecuteAddCommand(cmd)

	if err != nil {
		Log2Me("Error downloading magnet: " + err.Error())
		return
	}

	Log2Sender(q.Requester, "added torrent: "+add.Name)
	go q.WaitTillFinished(add)
}

func (q *DownloadQuery) WaitTillFinished(add transmission.TorrentAdded) {
	for {
		if TorrentIsFinished(add.ID) {
			Log2Sender(q.Requester, add.Name+" is finished")
			break
		}

		time.Sleep(time.Second * 5)
	}

	RemoveTorrent(add.ID)
}

// remove but keep files
func RemoveTorrent(id int) {
	cmd, _ := transmission.NewDelCmd(id, false)

	tclient.ExecuteCommand(cmd)
}

func TorrentIsFinished(id int) bool {
	torrents, err := tclient.GetTorrents()
	if err != nil {
		Log2Me(err.Error())
		return false
	}

	for _, torrent := range torrents {
		if torrent.ID == id {
			return torrent.PercentDone == float64(1)
		}
	}

	Log2Me(fmt.Sprintf("Cant find torrent by id %d", id))
	return true
}

func SetupTransmissionClient(settings Settings) {
	tclient = transmission.New(settings.TransmissionUrl, "", "")
}
