package bot

import (
    "github.com/dionbosschieter/downloader/searchprovider/rarbg"
    "github.com/dionbosschieter/downloader/searchprovider/thepiratebay"
)

var settings Settings

var allproviders = []SearchProvider{
    rarbg.SearchProvider{},
    thepiratebay.SearchProvider{},
}

// returns a searchprovider list sorted on the provided searchprovider names
func InitSearchProviders(providers []string) (searchproviders []SearchProvider) {
    searchproviders = make([]SearchProvider, len(allproviders))

    for counter,provider := range providers {
        for _,compareProvider := range allproviders {
            if compareProvider.Name() == provider {
                searchproviders[counter] = compareProvider
            }
        }
    }

    return searchproviders
}

func InitBot(settingsPath string) {
    if settings.FileExists(settingsPath) {
		settings.Parse(settingsPath)
	} else {
		panic("No settings.yaml is defined, see example.yaml")
	}

	searchproviders := InitSearchProviders(settings.SearchProviders)
	SetupTransmissionClient(settings)
	Log("Init downloader")
	SetupTalkyBot(settings, searchproviders)
}
