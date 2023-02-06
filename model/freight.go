package model

type Freight struct {
	Sku            string `json:"sku"`
	Stock          Stock  `json:"stock"`
	Weight         string `json:"weight"`
	PaidPrice      string `json:"paid_price"`
	IsTransferable bool   `json:"is_transferable"`
	ShipmentType   int    `json:"shipment_type"`
}

type Stock struct {
	Wharehouse int `json:"wharehouse"`
	Quantity   int `json:"quantity"`
}

var (
	ShipmentTypeMap = map[string]int{
		"NotFound":     0,
		"Own":          1,
		"Crossdocking": 2,
		"Consigned":    3,
		"DropShipping": 4,
		"Marketplace":  5,
	}
)
