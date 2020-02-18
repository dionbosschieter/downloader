package bot

import (
    "github.com/dionbosschieter/downloader/searchprovider"
    "gopkg.in/tucnak/telebot.v2"
	"time"
)

var tbot telebot.Bot

func SetupTalkyBot(settings Settings, providers []searchprovider.SearchProvider) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  settings.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}
	tbot = *bot

	tbot.Handle("/adds", func(m *telebot.Message) {
		if len(m.Payload) > 0 {
			query := Query{Title: m.Payload, Requester: m.Sender, Path: settings.SeriePath}
			query.Perform(providers, settings.SearchPostfixes)
		} else {
			tbot.Send(m.Sender, "Requires a payload /adds <payload>")
		}
	})

	tbot.Handle("/addm", func(m *telebot.Message) {
		if len(m.Payload) > 0 {
			query := Query{Title: m.Payload, Requester: m.Sender, Path: settings.MoviePath}
			query.Perform(providers, settings.SearchPostfixes)
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
