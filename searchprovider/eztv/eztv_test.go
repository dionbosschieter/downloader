package eztv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_buildSearchURL(t *testing.T) {
	assert.Equal(t, "https://eztv1.xyz/search/123", buildSearchURL("123"))
	assert.Equal(t, "https://eztv1.xyz/search/test-123", buildSearchURL("test 123"))
}

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	answer := provider.Search("test", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:def473f99aa4f0c99cd2b4ded600ab902e7bee0c")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	assert.Equal(t, "eztv", provider.Name())
}
