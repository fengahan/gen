package main

import (
	"github.com/gen/cmd/gen_route"
)

func main() {
	gen_route.Gen(".")
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run(":8082")
}
