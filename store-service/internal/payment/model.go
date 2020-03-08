package payment

type PaymentInformation struct {
	OrderID      int     `json:"order_id"`
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
	MerchantID   int     `json:"merchant_id"`
}

func NewShippingInfo(payment PaymentInformation) PaymentDetail {
	return PaymentDetail{
		CardNumber:   payment.CardNumber,
		CVV:          payment.CVV,
		ExpiredMonth: payment.ExpiredMonth,
		ExpiredYear:  payment.ExpiredYear,
		CardName:     payment.CardName,
		TotalPrice:   payment.TotalPrice,
	}
}
