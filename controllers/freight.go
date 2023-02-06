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

	chnFreight := make(chan []model.Freight)
	chnShipmentType := make(chan map[string]int)
	chnHigherPrice := make(chan model.Freight)

	go f.freightService.FillShipmentTypeCounts(chnFreight, chnShipmentType)
	chnFreight <- freight

	go f.freightService.HigherPrice(chnFreight, chnHigherPrice)
	chnFreight <- freight

	response := map[string]interface{}{
		"Shipment Type": <-chnShipmentType,
		"Higher Price":  <-chnHigherPrice,
	}

	c.JSON(http.StatusAccepted, &response)
}
