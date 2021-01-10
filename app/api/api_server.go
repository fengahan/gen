package main

import (
	"gen/app/api/gen_build"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gen_build.AmountRoute(gin.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello3", gen_build.InitRoleController().Hello3)
	r.Run(":8081")
}
