package controller

import (
	"net/http"
	"v1/domain"
	"v1/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService service.LoginService
}

func NewLoginController(loginService service.LoginService) *LoginController {
	return &LoginController{
		loginService: loginService,
	}
}

func (lc *LoginController) Handle(c *gin.Context) {
	var credentials domain.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := lc.loginService.AuthenticateUser(credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		Username: credentials.Username,
		Token:    tokenString,
	})
}
