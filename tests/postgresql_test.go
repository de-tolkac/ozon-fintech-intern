package tests

import (
	. "github.com/de-tolkac/ozon-fintech-intern/storage"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db PostgreSQL

func TestPostgresInit(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Equal(t, nil, err)

	err = db.Init()

	assert.Equal(t, nil, err)
}

/*
func TestPostgresFindEncodedUrl(t *testing.T) {
	var db PostgreSQL
	db.Init()

	for _, test := range tests {
		hash.codeToUrl[test.key] = test.value
	}

	for _, test := range tests {
		res, found := hash.FindEncodedUrl(test.key)
		assert.Equal(t, found, true)
		assert.Equal(t, test.value, res)
	}
}

func TestPostgresFindDecodedUrl(t *testing.T) {
	var hash HashTable
	hash.Init()

	for _, test := range tests {
		hash.urlToCode[test.value] = test.key
	}

	for _, test := range tests {
		res, found := hash.FindDecodedUrl(test.value)
		assert.Equal(t, found, true)
		assert.Equal(t, test.key, res)
	}
}

func TestPostgresSaveUrl(t *testing.T) {
	var hash HashTable
	hash.Init()

	for _, test := range tests {
		hash.SaveUrl(test.key, test.value)
	}

	for _, test := range tests {
		res, found := hash.FindEncodedUrl(test.key)
		assert.Equal(t, found, true)
		assert.Equal(t, test.value, res)

		res, found  = hash.FindDecodedUrl(test.value)
		assert.Equal(t, found, true)
		assert.Equal(t, test.key, res)
	}
}
*/
