package bot

import (
    "github.com/dionbosschieter/downloader/searchprovider/rarbg"
    "github.com/dionbosschieter/downloader/searchprovider/thepiratebay"
    "reflect"
)

var settings Settings

var allproviders = []SearchProvider{
    rarbg.SearchProvider{},
    thepiratebay.SearchProvider{},
}

func InArray(needle interface{}, haystack interface{}) bool {
    switch reflect.TypeOf(haystack).Kind() {
    case reflect.Slice, reflect.Array:
        list := reflect.ValueOf(haystack)

        for i:=0; i<list.Len();i++ {
            if reflect.DeepEqual(needle, list.Index(i).Interface()) {
                return true
            }
        }
    }

    return false
}

func InitSearchProviders(providers []string) (searchproviders []SearchProvider) {
    searchproviders = make([]SearchProvider, len(allproviders))

    for counter,provider := range allproviders {
        if InArray(provider.Name(), providers) {
            searchproviders[counter] = provider
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
