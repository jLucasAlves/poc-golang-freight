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

	response := fillShipmentTypeCounts(freight)

	c.JSON(http.StatusAccepted, &response)
}

func groupByShipmentType(freights []model.Freight) map[int]int {
	result := make(map[int]int)
	for _, freight := range freights {
		result[freight.ShipmentType]++
	}
	return result
}

func fillShipmentTypeCounts(freights []model.Freight) map[string]int {
	shipmentTypeMap := map[string]int{
		"NotFound":     0,
		"Own":          0,
		"Crossdocking": 0,
		"Consigned":    0,
		"DropShipping": 0,
		"Marketplace":  0,
	}
	groupedByShipmentType := groupByShipmentType(freights)

	for shipmentTypeName, shipmentType := range shipmentTypeMap {
		count, exists := groupedByShipmentType[shipmentType]
		if exists {
			shipmentTypeMap[shipmentTypeName] = count
		}
	}
	return shipmentTypeMap
}
