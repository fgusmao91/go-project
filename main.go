package main

import (
	"v1/runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers := runtime.InjectDependencies()

	r := gin.Default()

	r.POST("/login", controllers.LoginController.Login)
	r.POST("/register", controllers.LoginController.RegisterUser)
	r.POST("user/:user/add_authorization", controllers.LoginController.AddAuthorization)

	r.Run()
}