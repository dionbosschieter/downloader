package magnethair

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_buildSearchURL(t *testing.T) {
	assert.Equal(t, "https://magnetdl.hair/api.php?url=/q.php?q=123", buildSearchURL("123"))
	assert.Equal(t, "https://magnetdl.hair/api.php?url=/q.php?q=test+123", buildSearchURL("test 123"))
	assert.Equal(t, "https://magnetdl.hair/api.php?url=/q.php?q=Test", buildSearchURL("Test"))
}

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	answer := provider.Search("tttt", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:9949765AD15C921A2060557197EDCD21E6F1031B")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	assert.Equal(t, "magnethair", provider.Name())
}
