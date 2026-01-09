package bitmagnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	answer := provider.Search("Heroes S01", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:e4a0fbcdc4fd6e87c80e0a9fba01dd3b826b40b0")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	assert.Equal(t, "bitmagnet", provider.Name())
}
