package main

import (
	"gen/app/api/gen_build"
	"github.com/gin-gonic/gin"
)

func main() {
//	gen_route.Gen(".","app/api/gen_build/auto_gen_router.go")
	r := gen_build.AmountRoute(gin.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}
