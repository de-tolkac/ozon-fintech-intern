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