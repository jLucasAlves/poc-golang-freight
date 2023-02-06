package controllers

import (
	"net/http"

	"github.com/dafiti-group/poc-golang-freight/model"
	"github.com/dafiti-group/poc-golang-freight/service"
	"github.com/gin-gonic/gin"
)

type FreightController interface {
	FilterShipmentType(context *gin.Context)
}

type freightController struct {
	freightService service.FreightService
}

func NewFreightController(freightService service.FreightService) FreightController {
	return &freightController{
		freightService: freightService,
	}
}

func (f freightController) FilterShipmentType(c *gin.Context) {
	var freight []model.Freight

	if error := c.ShouldBindJSON(&freight); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	// cf := make(chan )

	response := f.freightService.FillShipmentTypeCounts(freight)

	c.JSON(http.StatusAccepted, &response)
}
