package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/domain"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("/login", func(c *gin.Context) {
		var credentials domain.Credentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "login")
	})

	r.Run()
}
