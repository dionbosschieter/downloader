package main

import (
    "github.com/dionbosschieter/downloader/bot"
    "flag"
)

// Allow users to provide a custom settings.yaml
func main() {
    settingsPath := flag.String("conf", "settings.yaml", "conf file probably settings.yaml")
    flag.Parse()

	bot.InitBot(*settingsPath)
}
