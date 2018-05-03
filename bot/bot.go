package bot

import (
    "github.com/dionbosschieter/downloader/searchprovider/rarbg"
    "github.com/dionbosschieter/downloader/searchprovider/thepiratebay"
    "reflect"
)

// var client piratebay.Piratebay
// var torrent piratebay.Torrent
// var rar *rarbg.Client
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
    counter := 0
    searchproviders = make([]SearchProvider, len(allproviders))

    for _,provider := range allproviders {
        counter++
        if InArray(provider, providers) {
            searchproviders[counter] = provider
        }
    }

    return searchproviders
}

func InitClient() {
	// client = piratebay.Piratebay{Url: "https://thepiratebay.org"}
	// rar, _ = rarbg.New(1337)
	// rar.Init()
}

func InitBot() {
    if settings.FileExists() {
		settings.Parse()
	} else {
		panic("No settings.yaml is defined, see example.yaml")
	}

	searchproviders := InitSearchProviders(settings.SearchProviders)
	SetupTransmissionClient(settings)
	Log("Init downloader")
	SetupTalkyBot(settings, searchproviders)
}
