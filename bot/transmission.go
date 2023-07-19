package bot

import (
	"fmt"
	"github.com/tubbebubbe/transmission"
	"log"
	"time"
)

var tclient transmission.TransmissionClient

func GetTorrents() string {
	returnstring := "Current torrents:\n"

	torrents, err := tclient.GetTorrents()
	if err != nil {
		return "Error listing current torrents: " + err.Error()
	}

	for _, t := range torrents {
		returnstring += t.Name + " " + fmt.Sprintf("%.4f", t.PercentDone*100) + "%\n"
	}

	return returnstring
}

func (q *Query) Download() {
	cmd, _ := transmission.NewAddCmdByMagnet(q.Magnet)
	cmd.SetDownloadDir(q.Path)
	add, err := tclient.ExecuteAddCommand(cmd)

	if err != nil {
		log.Println("Error downloading magnet: " + err.Error())
		return
	}

	q.Log2Requester("added torrent: " + add.Name)
	go q.WaitTillFinished(add)
}

func (q *Query) WaitTillFinished(add transmission.TorrentAdded) {
	for {
		if TorrentIsFinished(add.ID) {
			q.Log2Requester(add.Name + " is finished")
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

func ClearTorrents() (ret string) {
	ret = "Removed torrents:\n"
	torrents, err := tclient.GetTorrents()
	if err != nil {
		return "Error listing current torrents: " + err.Error()
	}

	for _, t := range torrents {
		cmd, _ := transmission.NewDelCmd(t.ID, true)

		ret += "Removing " + t.Name + "\n"
		_, _ = tclient.ExecuteCommand(cmd)
	}
	return
}

func TorrentIsFinished(id int) bool {
	torrents, err := tclient.GetTorrents()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	for _, torrent := range torrents {
		if torrent.ID == id {
			return torrent.PercentDone == float64(1)
		}
	}

	log.Println(fmt.Sprintf("Cant find torrent by id %d", id))
	return true
}

func SetupTransmissionClient(settings Settings) {
	tclient = transmission.New(settings.TransmissionUrl, "", "")
}
