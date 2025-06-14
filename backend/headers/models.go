package headers

type SKUItem struct {
	SKU int `json:"sku"`
	Qty int `json:"qty"`
}

type Pallet struct {
	ID      int       `json:"ID"`
	Aisle   int       `json:"aisle"`
	Bay     int       `json:"bay"`
	Side    string    `json:"side"`
	SKUList []SKUItem `json:"sku_list"`
}

type APIRequest struct {
	Action int    `json:"action"`
	Data   Pallet `json:"data"`
}
