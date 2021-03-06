package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	. "github.com/de-tolkac/ozon-fintech-intern/storage"
)

type Config struct {
	Storage   Storage
	UrlPrefix string
}

// Init config object
func (cfg *Config) Init(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		panic("Missed .env file")
	}

	// This prefix is concatenated with the short URL code
	cfg.UrlPrefix = os.Getenv("API_URL_PREFIX")

	// Parse command-line arguments
	storageType := flag.String("db", "postgresql", "Default: postgresql, HastTable-based: hash")
	flag.Parse()

	if *storageType == "postgresql" {
		cfg.Storage = new(PostgreSQL)
	} else if *storageType == "hash" {
		cfg.Storage = new(HashTable)
	} else {
		panic("Unknown storage type: " + *storageType + " (expected 'hash' or 'postgresql')")
	}

	// Init database object
	err = cfg.Storage.Init()
	if err != nil {
		panic("Error connecting and configuring the database: " + err.Error())
	}
}