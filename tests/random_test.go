package tests

import (
	. "github.com/de-tolkac/ozon-fintech-intern/random"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Well, it's a very difficult thing to test random generator
// In theory, we can provide statisticial analysys, but we won't ¯\_(ツ)_/¯
func TestRandomString(t *testing.T) {
	for i := 0; i < 100; i++ {
		assert.Equal(t, i, len(String(i)))
	}
}
