package main

import (
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

var tbot telebot.Bot
var me telebot.Chat

// todo: add queue commands
// todo: poll the queue
func SetupTalkyBot() {
	me = telebot.Chat{ID: settings.MasterChatId}
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  settings.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}
	tbot = *bot

	tbot.Send(&me, "# Bot is started")

	tbot.Handle("/adds", func(m *telebot.Message) {
		if len(m.Payload) > 0 {
			query := DownloadQuery{Title: m.Payload, Requester: m.Sender, Path: settings.SeriePath}
			query.Perform()
		} else {
			tbot.Send(m.Sender, "Requires a payload /adds <payload>")
		}
	})

	tbot.Handle("/addm", func(m *telebot.Message) {
		if len(m.Payload) > 0 {
			query := DownloadQuery{Title: m.Payload, Requester: m.Sender, Path: settings.MoviePath}
			query.Perform()
		} else {
			tbot.Send(m.Sender, "Requires a payload /addm <payload>")
		}
	})

	tbot.Handle("/status", func(m *telebot.Message) {
		tbot.Send(m.Sender, GetTorrents())
	})

	tbot.Handle("/help", func(m *telebot.Message) {
		tbot.Send(m.Sender, "/addm <search title> for movies\n/adds <search title> for series")
	})

	tbot.Start()
}
