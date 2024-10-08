package yts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	answer := provider.Search("test", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:2C7244EF6EEC682996F806685D494648C9339C8C")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	assert.Equal(t, "yts", provider.Name())
}
