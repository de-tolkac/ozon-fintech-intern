package storage

import "sync"

type HashTable struct {
	UrlToCode map[string]string // long URL to short URL
	CodeToUrl map[string]string // short URL to long URL
	Mutex     sync.Mutex        // mutex )))
}

// Init hash tables
func (table *HashTable) Init() error {
	table.UrlToCode = make(map[string]string)
	table.CodeToUrl = make(map[string]string)

	return nil
}

func (table *HashTable) FindEncodedUrl(url string) (res string, found bool) {
	res, found = table.CodeToUrl[url]
	return
}

func (table *HashTable) FindDecodedUrl(url string) (res string, found bool) {
	res, found = table.UrlToCode[url]
	return
}

func (table *HashTable) SaveUrl(decodedUrl, encodedUrl string) {
	// Save in both tables
	table.CodeToUrl[decodedUrl] = encodedUrl
	table.UrlToCode[encodedUrl] = decodedUrl
}

// Lock mutex
func (table *HashTable) Lock() {
	table.Mutex.Lock()
}

// Unlock mutex
func (table *HashTable) Unlock() {
	table.Mutex.Unlock()
}
