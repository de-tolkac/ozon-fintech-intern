package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "github.com/de-tolkac/ozon-fintech-intern/config"
	"github.com/de-tolkac/ozon-fintech-intern/random"
	"github.com/de-tolkac/ozon-fintech-intern/url"
)

type Request struct {
	Url string `json:"url" binding:"required"`
}

func Encode(cfg *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json Request
		if err := c.ShouldBindJSON(&json); err != nil {
			// Invalid request body
			c.JSON(http.StatusBadRequest, gin.H{
				"encodedUrl": "",
				"code":       2, // Bad Request
				"error":      "Invalid request body",
			})
			return
		}

		// Uncomment if you need to truncate trailing slashes from URL
		// urlString := url.Truncate(json.Url)
		urlString := json.Url
		if !url.Validate(urlString) {
			// Invalid URL in request
			c.JSON(http.StatusOK, gin.H{
				"encodedUrl": "",
				"code":       1, // Invalid URL
				"error":      "Invalid URL",
			})
			return
		}

		cfg.Storage.Lock()
		encodedUrl, found := cfg.Storage.FindEncodedUrl(urlString)
		if !found {
			// There is a small chance (N / 63^10, where N - number of existing keys)
			// that we will generate an already existing string.
			// Generating strings in a loop until we get a unique one.
			for {
				encodedUrl = random.String(10)
				_, alreadyExists := cfg.Storage.FindEncodedUrl(encodedUrl)
				if !alreadyExists {
					(*cfg).Storage.SaveUrl(urlString, encodedUrl)

					break
				}
			}
		}
		cfg.Storage.Unlock()

		c.JSON(http.StatusOK, gin.H{
			"encodedUrl": cfg.UrlPrefix + encodedUrl,
			"code":       0, // OK
			"error":      "",
		})
	}
}
