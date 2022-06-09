package tests

import (
	. "github.com/de-tolkac/ozon-fintech-intern/url"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	assert.Equal(t, true, Validate("http://url.ru"))
	assert.Equal(t, true, Validate("http://url.ru/"))
	assert.Equal(t, true, Validate("http://url.ru////"))
	assert.Equal(t, true, Validate("https://url.ru"))

	assert.Equal(t, false, Validate("://url.ru"))
	assert.Equal(t, false, Validate("url.ru"))
	assert.Equal(t, false, Validate("url.ru/"))
	assert.Equal(t, false, Validate(""))
	//assert.Equal(t, false, Validate("http://.ru"))
	//assert.Equal(t, false, Validate("http://url."))
}

/*
func TestTruncateSlashes(t *testing.T) {
	assert.Equal(t, "http://url.ru", TruncateSlashes("http://url.ru"))
	assert.Equal(t, "http://url.ru", TruncateSlashes("http://url.ru/"))
	assert.Equal(t, "http://url.ru", TruncateSlashes("http://url.ru////"))
	assert.Equal(t, "http://url.ru/a", TruncateSlashes("http://url.ru/a"))
	assert.Equal(t, "", TruncateSlashes(""))
	assert.Equal(t, "", TruncateSlashes("/"))
	assert.Equal(t, "", TruncateSlashes("/////"))
}
*/