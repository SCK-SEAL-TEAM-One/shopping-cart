package order

type OrderRepositoryMySQL struct {
}

func (orderRepository OrderRepositoryMySQL) CreateShipping(orderID int, submittedOrder SubmitedOrder) (int, error) {
	return 0, nil
}
