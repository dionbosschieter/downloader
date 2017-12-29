package main

import (
	"time"
    "gopkg.in/tucnak/telebot.v2"
)

var tbot telebot.Bot
var me = telebot.Chat{ID: 1337, }

func SetupTalkyBot() {
    bot, err := telebot.NewBot(telebot.Settings{
        Token:  "<token>",
        Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
    })
    if err != nil {
        panic(err)
    }
    tbot = *bot

    tbot.Send(&me, "# Bot is started")

    tbot.Handle("/adds", func(m *telebot.Message) {
        if len(m.Payload) > 0 {
            Search(m.Payload, "<TVSHOWPATH>")
        } else {
            tbot.Send(m.Sender, "Requires a payload /adds <payload>")
        }
    })

    tbot.Handle("/addm", func(m *telebot.Message) {
        if len(m.Payload) > 0 {
            Search(m.Payload, "<MOVIEPATH>")
        } else {
            tbot.Send(m.Sender, "Requires a payload /addm <payload>")
        }
    })

    tbot.Handle("/status", func(m *telebot.Message) {
        tbot.Send(m.Sender, GetTorrents())
    })

    tbot.Start()
}
