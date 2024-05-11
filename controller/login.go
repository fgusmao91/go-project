package controller

import (
	"net/http"
	"v1/domain"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (lc *LoginController) Handle(c *gin.Context) {
	var credentials domain.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, credentials)
}
