package tests

import (
	. "github.com/de-tolkac/ozon-fintech-intern/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var Cfg *Config
var ConfigInited bool

func InitConfig() {
	Cfg = new(Config)
	Cfg.Init("../testing.env")
	ConfigInited = true
}

func TestConfig(t *testing.T) {
	if !ConfigInited {
		InitConfig()
	}

	assert.Equal(t, os.Getenv("API_URL_PREFIX"), Cfg.UrlPrefix)
}
