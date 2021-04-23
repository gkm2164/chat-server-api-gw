package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func main() {
	server := gin.New()
	server.GET("/connect", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
			c.JSON(400, gin.H{
				"statusCode": 400,
				"body": "Bad request",
			})
		} else {
			fmt.Println(string(body))
			c.JSON(200, gin.H{
				"statusCode": 200,
				"body":       "Connected.",
			})
		}
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
	if err := server.Run(":8081"); err != nil {
		log.Fatalf("error while loading server: %v", err)
	}
}