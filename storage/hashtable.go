package storage

import "sync"

type HashTable struct {
	urlToCode map[string]string
	codeToUrl map[string]string
	mutex     sync.Mutex
}

func (table *HashTable) Init() error {
	table.urlToCode = make(map[string]string)
	table.codeToUrl = make(map[string]string)

	return nil
}

func (table *HashTable) FindEncodedUrl(url string) (res string, found bool) {
	res, found = table.codeToUrl[url]
	return
}

func (table *HashTable) FindDecodedUrl(url string) (res string, found bool) {
	res, found = table.urlToCode[url]
	return
}

func (table *HashTable) SaveUrl(decodedUrl, encodedUrl string) {
	table.codeToUrl[decodedUrl] = encodedUrl
	table.urlToCode[encodedUrl] = decodedUrl
}

func (table *HashTable) Lock() {
	table.mutex.Lock()
}

func (table *HashTable) Unlock() {
	table.mutex.Unlock()
}
