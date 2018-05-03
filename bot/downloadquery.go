package bot

import (
	"gopkg.in/tucnak/telebot.v2"
)

type DownloadQuery struct {
	Title     string
	Requester *telebot.User
	Path      string
	Magnet    string
	IsMovie   bool
}
