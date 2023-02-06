package service

import (
	"github.com/dafiti-group/poc-golang-freight/model"
)

type FreightService interface {
	GroupByShipmentType(freights []model.Freight) map[int]int
	FillShipmentTypeCounts(freights []model.Freight) map[string]int
}

type freightService struct{}

func NewFreightService() FreightService {
	return &freightService{}
}

func (service *freightService) GroupByShipmentType(freights []model.Freight) map[int]int {
	result := make(map[int]int)
	for _, freight := range freights {
		result[freight.ShipmentType]++
	}
	return result
}

func (service *freightService) FillShipmentTypeCounts(freights []model.Freight) map[string]int {
	result := make(map[string]int)

	groupedByShipmentType := service.GroupByShipmentType(freights)

	for shipmentTypeName, shipmentType := range model.ShipmentTypeMap {
		count, exists := groupedByShipmentType[shipmentType]
		if exists {
			result[shipmentTypeName] = count
		} else {
			result[shipmentTypeName] = 0
		}
	}
	return result
}
