package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.POST("/connect", func(c *gin.Context) {
		var x gin.H
		_ = c.BindJSON(&x)

		fmt.Println(x)

		c.JSON(200, gin.H{
			"statusCode": 200,
			"body": "Connected.",
		})
	})

	server.POST("/sendmessage", func(c *gin.Context) {
		var x gin.H
		_ = c.BindJSON(&x)

		fmt.Println(x)

		c.JSON(200, gin.H{
			"statusCode": 200,
			"body": "Connected.",
		})
	})

	server.POST("/disconnect", func(c *gin.Context) {
		var x gin.H
		_ = c.BindJSON(&x)

		fmt.Println(x)

		c.JSON(200, gin.H{
			"statusCode": 200,
			"body": "Connected.",
		})
	})
}