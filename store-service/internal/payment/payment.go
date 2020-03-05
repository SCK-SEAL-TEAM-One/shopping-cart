package payment

type PaymentInterface interface {
	ConfirmPayment(paymentdetail PaymentDetail) string
}

type PaymentService struct {
	PaymentRepository PaymentRepository
}

func (paymentService PaymentService) ConfirmPayment(paymentdetail PaymentDetail) string {
	return "วันเวลาที่ชำระเงิน 1/3/2563 13:30:00 หมายเลขคำสั่งซื้อ 8004359103 คุณสามารถติดตามสินค้าผ่านช่องทาง Kerry หมายเลข Tracking 1785261900"
}
