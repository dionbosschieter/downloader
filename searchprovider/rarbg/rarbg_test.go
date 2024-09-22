package rarbg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}

	answer := provider.Search("test", []string{})

	assert.Contains(t, "magnet:?xt=urn:btih:BDFED52F8B46295DEB23FEA529D00C4FC9DAE8FA", answer)
}

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}

	assert.Equal(t, "rarbg", provider.Name())
}
