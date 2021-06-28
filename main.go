package main

import (
    "flag"
    "github.com/dionbosschieter/downloader/bot"
)

// Allow users to provide a custom settings.yaml
func main() {
    settingsPath := flag.String("conf", "settings.yaml", "conf file probably settings.yaml")
    flag.Parse()

    downloader := bot.Bot{
        SettingsPath: *settingsPath,
    }

    downloader.Start()
}
