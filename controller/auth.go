package controller

import (
	"net/http"
	"v1/service"
	"v1/dto"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (lc *AuthController) Login(c *gin.Context) {
	var credentials dto.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := lc.authService.AuthenticateUser(credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Username: credentials.Username,
		Token:    tokenString,
	})
}

func (lc *AuthController) RegisterUser(c *gin.Context) {
	var credentials dto.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := lc.authService.RegisterUser(credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "user registered successfully")
}

func (lc *AuthController) AddAuthorization(c *gin.Context) {
	var credentials dto.Authorization
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userName := c.Param("user")

	err := lc.authService.AddAuthorization(userName, credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, "authorization added successfully")
}