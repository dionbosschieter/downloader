package rarbg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}

	provider.Init()

	answer := provider.Search("test", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:5a547a8585d4dd88e884e33326bb40f671a2b3a7&dn=The.Greatest.Beer.Run.Ever.2022.1080p.WEB.h264-TRUFFLE")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}

	assert.Equal(t, "rarbg", provider.Name())
}
