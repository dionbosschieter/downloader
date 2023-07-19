package bot

import (
	"fmt"
	"github.com/dionbosschieter/downloader/searchprovider"
	"github.com/dionbosschieter/downloader/searchprovider/leet"
	"github.com/dionbosschieter/downloader/searchprovider/magnet"
	"github.com/dionbosschieter/downloader/searchprovider/magnetdl"
	"github.com/dionbosschieter/downloader/searchprovider/yts"
	"log"
)

// Bot is the entry point of the bot logic, see Start()
type Bot struct {
	// SettingsPath is the location to the settings file to read
	SettingsPath string
	// settings contains all the information we need to run this bot
	settings Settings
}

func (b *Bot) Start() {
	if b.settings.FileExists(b.SettingsPath) {
		b.settings.Parse(b.SettingsPath)
	} else {
		fmt.Println("No settings.yaml is defined, see example.yaml")
		return
	}

	providers := InitSearchProviders(b.settings.SearchProviders)
	SetupTransmissionClient(b.settings)
	log.Println("Init downloader")

	RunTelegramBot(b.settings, providers)
}

// InitSearchProviders returns a SearchProvider list sorted on the provided provider names
func InitSearchProviders(providers []string) []searchprovider.SearchProvider {
	searchProviders := make([]searchprovider.SearchProvider, len(providers))

	count := 0
	for _, providerName := range providers {
		provider := getProviderByName(providerName)

		if provider == nil {
			log.Printf("Can't find given search provider '%s'\n", providerName)
		}

		if provider.Name() == providerName {
			provider.Init()
			searchProviders[count] = provider
			count++
		}
	}

	return searchProviders
}

// getProviderByName returns a provider or nil if there no match can be found
func getProviderByName(providerName string) searchprovider.SearchProvider {
	if providerName == "1337x" {
		return &leet.SearchProvider{}
	}
	if providerName == "yts" {
		return &yts.SearchProvider{}
	}
	if providerName == "magnet" {
		return &magnet.SearchProvider{}
	}
	if providerName == "magnetdl" {
		return &magnetdl.SearchProvider{}
	}

	return nil
}
