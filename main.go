package main

import (
	"net/http"
	"v1/runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers := runtime.InjectDependencies()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })
	r.POST("/login", controllers.LoginController.Handle)

	r.Run()
}