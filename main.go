package main

import "github.com/gin-gonic/gin"

func main() {
	start()
}

func start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.GET("/tracks/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id":     c.Params.ByName("id"),
			"name":   "Suzuka",
			"length": "1400",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
