package cart

type SubmitedCart struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	ID        int `json:"id" db:"id"`
	UserID    int `json:"user_id" db:"user_id"`
	ProductID int `json:"product_id" db:"product_id"`
	Quantity  int `json:"quantity" db:"quantity"`
}

type CartDetail struct {
	ID        int     `json:"id" db:"id"`
	UserID    int     `json:"user_id" db:"user_id"`
	ProductID int     `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Name      string  `json:"product_name" db:"product_name"`
	Price     float64 `json:"product_price" db:"product_price"`
	Image     string  `json:"product_image" db:"image_url"`
	Stock     int     `json:"stock" db:"stock"`
	Brand     string  `json:"product_brand" db:"product_brand"`
}
