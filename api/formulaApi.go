package api

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
