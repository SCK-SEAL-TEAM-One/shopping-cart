package api

type SubmmitedOrder struct {
	Cart                 []OrderProduct `json:"cart"`
	ShippingMethod       int            `json:"shipping_method"`
	ShippingAddress      string         `json:"shipping_address"`
	ShippingSubDistrict  string         `json:"shipping_sub_disterict"`
	ShippingDistrict     string         `json:"shipping_district"`
	ShippingProvince     string         `json:"shipping_province"`
	ShippingZipCode      string         `json:"shipping_zip_code"`
	RecipientName        string         `json:"recipient_name"`
	RecipientPhoneNumber string         `json:"recipient_phone_number"`
}

type OrderProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderConfirmation struct {
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}
