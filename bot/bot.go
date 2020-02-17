package bot

import (
    "plugin"
)

var settings Settings

// returns a searchprovider list sorted on the provided searchprovider names
func InitSearchProviders(providers []string) []SearchProvider {
    searchProviders := make([]SearchProvider, len(providers))

    var count = 0
    for _,provider := range providers {
        plug, err := plugin.Open("./" + provider + ".so")
        if err != nil {
            Log2Me("Could not find plugin for provider " + provider)
            continue
        }

        symSearchProvider,err := plug.Lookup("SearchProvider")
        if err != nil {
            Log2Me("Could not find SearchProvider symbol for " + provider)
            continue
        }

        searchprovider,ok := symSearchProvider.(SearchProvider)
        if !ok {
            Log2Me("Unexpected type from SearchProvider: " + provider)
        }

        if searchprovider.Name() == provider {
            searchprovider.Init()
            searchProviders[count] = searchprovider
            count++
        }
    }

    return searchProviders
}

func InitBot(settingsPath string) {
    if settings.FileExists(settingsPath) {
		settings.Parse(settingsPath)
	} else {
		panic("No settings.yaml is defined, see example.yaml")
	}

	providers := InitSearchProviders(settings.SearchProviders)
    SetupTransmissionClient(settings)
	Log("Init downloader")

    SetupTalkyBot(settings, providers)
}
