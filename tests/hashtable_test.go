package tests

import (
	. "github.com/de-tolkac/ozon-fintech-intern/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

type pair struct {
	key   string
	value string
}

var validTests = []pair{
	pair{"code", "url-1"},
	pair{"another-code", "url-2"},
	pair{"string", "string"},
	pair{"yet-another-string", "yet-another-url"},
}

var invalidTests = []pair{
	pair{"url-1", "code"},
	pair{"url-2", "another-code"},
	pair{"string2", "string2"},
}

func TestHashInit(t *testing.T) {
	var hash HashTable
	err := hash.Init()
	assert.Equal(t, nil, err)
}

func TestHashFindEncodedUrl(t *testing.T) {
	var hash HashTable
	hash.Init()

	for _, test := range validTests {
		hash.CodeToUrl[test.key] = test.value
	}

	for _, test := range validTests {
		res, found := hash.FindEncodedUrl(test.key)

		assert.Equal(t, true, found)
		assert.Equal(t, test.value, res)
	}

	for _, test := range invalidTests {
		res, found := hash.FindEncodedUrl(test.key)

		assert.Equal(t, false, found)
		assert.NotEqual(t, test.value, res)
	}
}

func TestHashFindDecodedUrl(t *testing.T) {
	var hash HashTable
	hash.Init()

	for _, test := range validTests {
		hash.UrlToCode[test.value] = test.key
	}

	for _, test := range validTests {
		res, found := hash.FindDecodedUrl(test.value)

		assert.Equal(t, true, found)
		assert.Equal(t, test.key, res)
	}

	for _, test := range invalidTests {
		res, found := hash.FindDecodedUrl(test.value)

		assert.Equal(t, false, found)
		assert.NotEqual(t, test.key, res)
	}
}

func TestHashSaveUrl(t *testing.T) {
	var hash HashTable
	hash.Init()

	for _, test := range validTests {
		hash.SaveUrl(test.key, test.value)
	}

	for _, test := range validTests {
		res, found := hash.FindEncodedUrl(test.key)

		assert.Equal(t, found, true)
		assert.Equal(t, test.value, res)

		res, found = hash.FindDecodedUrl(test.value)

		assert.Equal(t, found, true)
		assert.Equal(t, test.key, res)
	}
}
