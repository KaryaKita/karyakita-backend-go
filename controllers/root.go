package controllers

import (
	"github.com/gin-gonic/gin"
)

type RootController struct{}

func (rootController RootController) GetRoot(c *gin.Context) {
	c.JSON(200,
		gin.H{"message": "Selamat datang di aplikasi karyakita.com"})
	return
}
