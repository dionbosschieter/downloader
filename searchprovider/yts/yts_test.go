package yts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	answer := provider.Search("test", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:353E43CACDB811FDED4E25C2C5134DF41E666489")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	assert.Equal(t, "yts", provider.Name())
}
