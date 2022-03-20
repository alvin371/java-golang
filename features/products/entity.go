package products

import "time"

type Products struct {
	ID            int
	NameID        string
	DescriptionID string
	NameEN        string
	DescriptionEN string
	Category      string
	Image         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Bussiness interface {
	GetAllProducts(Products) (prd []Products, err error)
	GetProductsById(id int) (prd Products, err error)
	CreateProducts(data Products) (err error)
	EditProducts(id int, data Products) (prd Products, err error)
	DeleteProducts(id int) (prd Products, err error)
}
type Data interface {
	SelectAllProducts(Products) (prd []Products, err error)
	SelectProductsById(id int) (prd Products, err error)
	InsertProducts(data Products) (err error)
	UpdateProducts(id int, data Products) (prd Products, err error)
	DestroyProducts(id int) (prd Products, err error)
}
