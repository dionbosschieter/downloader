package yts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	answer := provider.Search("test", []string{})

	assert.Contains(t, answer, "magnet:?xt=urn:btih:109F7EB71F1C2CDFF56DFA75385E98CD2A636A6D")
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	provider.Init()

	assert.Equal(t, "yts", provider.Name())
}
