package leet

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestSearchProvider_Search(t *testing.T) {
    provider := SearchProvider{}

    answer := provider.Search("test", []string{})

    assert.Equal(t, "magnet:?xt=urn:btih:7D056C4CC36D897F36477CEA876C9B137B7A2E38&dn=TittyAttack+21+01+02+Alyx+Star+Test+Them+Out++480p+MP4-XXX&tr=udp%3A%2F%2F9.rarbg.to%3A2940%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2750%2Fannounce&tr=udp%3A%2F%2Ftracker.trackerfix.com%3A83%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.pirateparty.gr%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2F&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce", answer)
}

func TestSearchProvider_Name(t *testing.T) {
    provider := SearchProvider{}

    assert.Equal(t, "1337x", provider.Name())
}
