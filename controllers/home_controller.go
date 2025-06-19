package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Go + Gin",
	})
}
