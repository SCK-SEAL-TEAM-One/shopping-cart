package product

type ProductDetail struct {
	ID       int     `json:"id"`
	Name     string  `json:"product_name"`
	Price    float64 `json:"product_price"`
	Image    string  `json:"product_image"`
	Quantity int     `json:"quantity"`
	Brand    string  `json:"product_brand"`
}
