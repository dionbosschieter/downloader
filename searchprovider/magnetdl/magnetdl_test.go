package magnetdl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_buildSearchURL(t *testing.T) {
	assert.Equal(t, "https://www.magnetdl.com/1/123/", buildSearchURL("123"))
	assert.Equal(t, "https://www.magnetdl.com/t/test-123/", buildSearchURL("test 123"))
	assert.Equal(t, "https://www.magnetdl.com/t/test/", buildSearchURL("Test"))
}

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	answer := provider.Search("tttt", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:97c2b6467f0e180d0793f43ab42a908bbc9c2603")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	assert.Equal(t, "magnetdl", provider.Name())
}
