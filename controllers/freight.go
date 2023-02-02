package controllers

import (
	"net/http"

	"github.com/dafiti-group/poc-golang-freight/model"
	"github.com/gin-gonic/gin"
)

func Freight(c *gin.Context) {
	var freight []model.Freight

	if error := c.ShouldBindJSON(&freight); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusAccepted, &freight)
}
