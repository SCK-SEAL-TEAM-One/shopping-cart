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

type SubmitedOrder struct {
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

type Order struct {
	OrderID    int
	TotalPrice float64
}

func (s SubmitedOrder) GetShippingFee() float64 {
	return 2.00
}
