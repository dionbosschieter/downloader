package bot

import (
	"github.com/dionbosschieter/downloader/searchprovider"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var tbot telebot.Bot

func RunTelegramBot(settings Settings, providers []searchprovider.SearchProvider) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  settings.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Println("Problem creating telegram bot:" + err.Error())
		return
	}
	tbot = *bot

	tbot.Handle("/adds", func(m *telebot.Message) {
		if len(m.Payload) > 0 {
			query := Query{Title: m.Payload, Requester: m.Sender, Path: settings.SeriePath}
			query.Perform(providers, settings.SearchPostfixes)
		} else {
			_, _ = tbot.Send(m.Sender, "Requires a payload /adds <payload>")
		}
	})

	tbot.Handle("/addm", func(m *telebot.Message) {
		if len(m.Payload) > 0 {
			query := Query{Title: m.Payload, Requester: m.Sender, Path: settings.MoviePath}
			query.Perform(providers, settings.SearchPostfixes)
		} else {
			_, _ = tbot.Send(m.Sender, "Requires a payload /addm <payload>")
		}
	})

	tbot.Handle("/status", func(m *telebot.Message) {
		_, _ = tbot.Send(m.Sender, GetTorrents())
	})

	tbot.Handle("/help", func(m *telebot.Message) {
		_, _ = tbot.Send(m.Sender, "/addm <search title> for movies\n/adds <search title> for series\n/status\n/clear")
	})

	tbot.Handle("/clear", func(m *telebot.Message) {
		_, _ = tbot.Send(m.Sender, ClearTorrents())
	})

	tbot.Start()
}
