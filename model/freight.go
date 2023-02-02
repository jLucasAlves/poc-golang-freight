package model

type Freight struct {
	Sku            string       `json:"sku"`
	Stock          Stock        `json:"stock"`
	Weight         string       `json:"weight"`
	PaidPrice      string       `json:"paid_price"`
	IsTransferable bool         `json:"is_transferable"`
	ShipmentType   ShipmentType `json:"shipment_type"`
}

type Stock struct {
	Wharehouse int `json:"wharehouse"`
	Quantity   int `json:"quantity"`
}

type ShipmentType int

const (
	NotFound ShipmentType = iota
	Own
	Crossdocking
	Consigned
	DropShipping
	Marketplace
)
