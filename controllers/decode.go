package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "github.com/de-tolkac/ozon-fintech-intern/config"
)

func Decode(cfg *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedUrl, found := cfg.Storage.FindDecodedUrl(c.Param("short-url"))
		if !found {
			c.JSON(http.StatusOK, gin.H{
				"decodedUrl": "",
				"code":       1, // Url not found in DB
				"error":      "Url not found",
			})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"decodedUrl": decodedUrl,
			"code":       0, // Success
			"error":      "",
		})
	}
}
