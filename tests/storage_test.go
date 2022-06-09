package tests

import (
	. "github.com/de-tolkac/ozon-fintech-intern/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Check if database structures implements Storage interface
func TestInterfaces(t *testing.T) {
	hash := new(HashTable)
	_, hashFits := interface{}(hash).(Storage)
	assert.Equal(t, true, hashFits)

	postgres := new(PostgreSQL)
	_, postgresFits := interface{}(postgres).(Storage)
	assert.Equal(t, true, postgresFits)
}
