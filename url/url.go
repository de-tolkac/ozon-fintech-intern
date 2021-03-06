package url

import "net/url"

func Validate(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	return err == nil
}

/* 
// Well, smh trailing slash matters for URLs, so we need to disable this function
func Truncate(url string) string {
	// First comprassion prevents SEGFAULT when empty string passed
	if len(url) == 0 || url[len(url)-1] != '/' {
		return url
	}

	idx := len(url) - 1
	for ; idx >= 0 && (url[idx] == '/' || url[idx] == ' '); idx-- {
	} // C-style stuff....

	return url[:idx+1]
}
*/