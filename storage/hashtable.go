package storage

import "sync"

type HashTable struct {
	UrlToCode map[string]string
	CodeToUrl map[string]string
	Mutex     sync.Mutex
}

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
	table.CodeToUrl[decodedUrl] = encodedUrl
	table.UrlToCode[encodedUrl] = decodedUrl
}

func (table *HashTable) Lock() {
	table.Mutex.Lock()
}

func (table *HashTable) Unlock() {
	table.Mutex.Unlock()
}
