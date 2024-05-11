package main

import (
	"net/http"
	"v1/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controllers := runtime()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("/login", controllers.loginController.Handle)

	r.Run()
}

type controllers struct {
	loginController *controller.LoginController
}

func runtime() controllers {
	loginController := controller.NewLoginController()

	return controllers{
		loginController: loginController,
	}
}
