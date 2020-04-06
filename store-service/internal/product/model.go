package product

type ProductResult struct {
	Total    int       `json:"total"`
	Products []Product `json:"products"`
}

type Product struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"product_name" db:"product_name"`
	Price float64 `json:"product_price" db:"product_price"`
	Image string  `json:"product_image" db:"image_url"`
}

type ProductDetail struct {
	ID       int     `json:"id" db:"id"`
	Name     string  `json:"product_name" db:"product_name"`
	Price    float64 `json:"product_price" db:"product_price"`
	Image    string  `json:"product_image" db:"image_url"`
	Quantity int     `json:"quantity" db:"quantity"`
	Brand    string  `json:"product_brand" db:"product_brand"`
}
