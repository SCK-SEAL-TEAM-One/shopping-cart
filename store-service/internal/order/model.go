package order

type ShippingInfo struct {
	ShippingMethod       int
	ShippingAddress      string
	ShippingSubDistrict  string
	ShippingDistrict     string
	ShippingProvince     string
	ShippingZipCode      string
	RecipientName        string
	RecipientPhoneNumber string
}

type OrderProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
