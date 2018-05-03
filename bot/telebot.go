package bot

import (
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

var tbot telebot.Bot
var me telebot.Chat

func (q *DownloadQuery) Perform(searchproviders []SearchProvider, searchpostfixes []string) {
    for _,provider := range searchproviders {
        Log("Searching for " + q.Title + " with provider " + provider.Name())
        q.Magnet = provider.Search(q.Title, searchpostfixes)

        if q.Magnet != "" {
            q.Download()
            break
        }
    }

	// var rarbgQuery = map[string]string{"search_string": q.Title, "sort": "seeders", "category": "tv"}
	// if q.IsMovie {
	// 	rarbgQuery["category"] = "movies"
	// }
	// pbytorrents, pbyerr := client.Search(q.Title)
	// rarresult, rarerr := rar.Search(rarbgQuery)
    //
	// if pbyerr != nil {
	// 	Log2Me(pbyerr.Error())
	// }
	// if rarerr != nil {
	// 	Log2Me(rarerr.Error())
	// 	return
	// }
    //
	// if len(pbytorrents) == 0 && len(rarresult.Torrents) == 0 {
	// 	Log2Sender(q.Requester, "Could not find any result for: "+q.Title)
	// 	return
	// }
    //
	// if len(rarresult.Torrents) > 0 {
	// 	q.Magnet = rarresult.Torrents[0].MagnetURL
	// 	q.Download()
	// } else {
	// 	q.Magnet = pbytorrents[0].MagnetLink
	// 	q.Download()
	// }
}

// todo: add queue commands
// todo: poll the queue
func SetupTalkyBot(settings Settings, searchproviders []SearchProvider) {
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
			query.Perform(searchproviders, settings.SearchPostfixes)
		} else {
			tbot.Send(m.Sender, "Requires a payload /adds <payload>")
		}
	})

	tbot.Handle("/addm", func(m *telebot.Message) {
		if len(m.Payload) > 0 {
			query := DownloadQuery{Title: m.Payload, Requester: m.Sender, Path: settings.MoviePath}
			query.Perform(searchproviders, settings.SearchPostfixes)
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
