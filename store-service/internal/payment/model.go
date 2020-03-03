package payment

type PaymentInformation struct {
	PaymentType  string  `json:"payment_type"`
	Type         string  `json:"type"`
	CardNumber   string  `json:"card_number"`
	CVV          string  `json:"cvv"`
	ExpiredMonth int     `json:"expired_month"`
	ExpiredYear  int     `json:"expired_year"`
	CardName     string  `json:"card_name"`
	TotalPrice   float64 `json:"total_price"`
}

type PaymentDetail struct {
	CardNumber   string  `json:"card_number"`
	CVV          string  `json:"cvv"`
	ExpiredMonth int     `json:"expired_month"`
	ExpiredYear  int     `json:"expired_year"`
	CardName     string  `json:"card_name"`
	TotalPrice   float64 `json:"total_price"`
	MerchantID   int     `json:"MerchantID"`
}
