package product

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"product_price"`
	Image string  `json:"product_image"`
}
