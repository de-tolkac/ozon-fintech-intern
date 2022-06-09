package storage

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	connection *sql.DB
	mutex      sync.Mutex
}

func (db *PostgreSQL) Init() (err error) {
	// Generate url for PostgreSQL connection
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRESQL_USER"), 
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_DB_NAME"))

	db.connection, err = sql.Open("postgres", connStr)
	if err == nil {
		// We successfully connected to DB, now creating table (if not exists) -- migration of smoker
		_, err = db.connection.Exec(`CREATE TABLE IF NOT EXISTS url (
								id serial PRIMARY KEY, 
								long_url VARCHAR(2048) UNIQUE NOT NULL,
								short_code VARCHAR(10) UNIQUE NOT NULL)`)
	}

	return err
}

func (db *PostgreSQL) FindEncodedUrl(url string) (res string, found bool) {
	err := db.connection.QueryRow("SELECT short_code FROM url WHERE long_url = $1", url).Scan(&res)
	if err != nil {
		return
	}
	found = true

	return
}

func (db *PostgreSQL) FindDecodedUrl(code string) (res string, found bool) {
	err := db.connection.QueryRow("SELECT long_url FROM url WHERE short_code = $1", code).Scan(&res)
	if err != nil {
		return
	}
	found = true

	return
}

func (db *PostgreSQL) SaveUrl(decodedUrl, encodedUrl string) {
	db.connection.QueryRow("INSERT INTO url (long_url, short_code) VALUES ($1, $2)", decodedUrl, encodedUrl)
}

func (db *PostgreSQL) Lock() {
	db.mutex.Lock()
}

func (db *PostgreSQL) Unlock() {
	db.mutex.Unlock()
}
