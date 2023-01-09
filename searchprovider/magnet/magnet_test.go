package magnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProvider_Name(t *testing.T) {
	provider := SearchProvider{}
	assert.Equal(t, "magnet", provider.Name())
}

func TestSearchProvider_Search(t *testing.T) {
	provider := SearchProvider{}
	assert.Equal(t, "magnet:?", provider.Search("magnet:?", []string{}))
}
