package storage

type Storage interface {
	Init() error
	FindEncodedUrl(url string) (string, bool)
	FindDecodedUrl(url string) (string, bool)
	SaveUrl(decodedUrl, encodedUrl string)
	Lock()
	Unlock()
}
