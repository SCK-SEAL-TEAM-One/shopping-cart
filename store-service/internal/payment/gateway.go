package payment

type BankGateway struct {
	BankEndpoint string
}

func (gateway BankGateway) Payment(paymentDetail PaymentDetail) (string, error) {
	return "", nil
}
