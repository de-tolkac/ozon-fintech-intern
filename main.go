package main

import (
	"github.com/gin-gonic/gin"

	. "github.com/de-tolkac/ozon-fintech-intern/config"
	"github.com/de-tolkac/ozon-fintech-intern/controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfg := new(Config)
	cfg.Init(".env")

	r := gin.Default()

	r.POST("/encode", controllers.Encode(cfg))
	r.GET("/:short-url", controllers.Decode(cfg))

	r.Run()
}
