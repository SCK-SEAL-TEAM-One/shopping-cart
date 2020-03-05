package order

type SubmitedOrder struct {
	Cart                 []OrderProduct `json:"cart"`
	ShippingMethod       string         `json:"shipping_method"`
	ShippingAddress      string         `json:"shipping_address"`
	ShippingSubDistrict  string         `json:"shipping_sub_disterict"`
	ShippingDistrict     string         `json:"shipping_district"`
	ShippingProvince     string         `json:"shipping_province"`
	ShippingZipCode      string         `json:"shipping_zip_code"`
	RecipientName        string         `json:"recipient_name"`
	RecipientPhoneNumber string         `json:"recipient_phone_number"`
}

type ShippingInfo struct {
	ShippingMethod       string `db:"method"`
	ShippingAddress      string `db:"address"`
	ShippingSubDistrict  string `db:"sub_district"`
	ShippingDistrict     string `db:"district"`
	ShippingProvince     string `db:"province"`
	ShippingZipCode      string `db:"zip_code"`
	RecipientName        string `db:"recipient"`
	RecipientPhoneNumber string `db:"phone_number"`
}

type OrderProduct struct {
	ProductID int `json:"product_id" db:"product_id"`
	Quantity  int `json:"quantity" db:"quantity"`
}

type Order struct {
	OrderID    int
	TotalPrice float64
}

func (s SubmitedOrder) GetShippingFee() float64 {
	return 2.00
}

func (s SubmitedOrder) GetShippingMethodProvider() string {
	return "Kerry"
}
