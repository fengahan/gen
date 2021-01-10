package main

import (
	"gen/app/api/gen_build"

	"gen/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.AmountConfig("config/application.yaml")

	r := gen_build.AmountRoute(gin.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}
