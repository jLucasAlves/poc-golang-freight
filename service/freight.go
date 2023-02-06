package service

import (
	"strconv"
	"strings"

	"github.com/dafiti-group/poc-golang-freight/model"
)

type FreightService interface {
	FillShipmentTypeCounts(chnFreight chan []model.Freight, chnShipmentType chan map[string]int)
	HigherPrice(chnFreight chan []model.Freight, chnHigherPrice chan model.Freight)
}

type freightService struct{}

func NewFreightService() FreightService {
	return &freightService{}
}

func (service *freightService) FillShipmentTypeCounts(chnFreight chan []model.Freight, chnShipmentType chan map[string]int) {
	freight := <-chnFreight
	shipmentTypeCounts := make(map[string]int)

	for _, f := range freight {
		for shipmentTypeName, id := range model.ShipmentTypeMap {
			if f.ShipmentType == id {
				shipmentTypeCounts[shipmentTypeName]++
			}
		}
	}

	chnShipmentType <- shipmentTypeCounts
}

func (service *freightService) HigherPrice(chnFreight chan []model.Freight, chnHigherPrice chan model.Freight) {
	var freight model.Freight
	var currentHighestPrice float64

	for freights := range chnFreight {
		for _, freightItem := range freights {
			price, _ := strconv.ParseFloat(strings.TrimLeft(freightItem.PaidPrice, "$"), 64)
			if currentHighestPrice == 0 || price > currentHighestPrice {
				currentHighestPrice = price
				freight = freightItem
			}
		}
		chnHigherPrice <- freight
	}
}
