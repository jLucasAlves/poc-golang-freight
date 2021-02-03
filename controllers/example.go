package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Example is an example of controller
func Example(c *gin.Context) {
	logrus.Info("estou aqui")
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Say helo for your new Golang+gin microservise (By Dafiti-group)",
	})
}
