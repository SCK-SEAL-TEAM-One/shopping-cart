package product

type ProductRepository interface {
	GetProductByID(ID int) (Product, error)
}
