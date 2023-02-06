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
	shipmentTypeMap := map[string]int{
		"NotFound":     0,
		"Own":          1,
		"Crossdocking": 2,
		"Consigned":    3,
		"DropShipping": 4,
		"Marketplace":  5,
	}
	groupedByShipmentType := service.GroupByShipmentType(freights)

	for shipmentTypeName, shipmentType := range shipmentTypeMap {
		count, exists := groupedByShipmentType[shipmentType]
		if exists {
			shipmentTypeMap[shipmentTypeName] = count
		}
	}
	return shipmentTypeMap
}
