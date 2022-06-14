package controllers

import (
	"github.com/gin-gonic/gin"
)

// Example is an example of controller
func Example(c *gin.Context) {
	// logrus.Info("estou aqui")
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Say hello for your new Golang+gin microservise template-golang (By Dafiti-group)",
	})
}
