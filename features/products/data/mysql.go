package data

import (
	"errors"
	"fmt"
	"java-agro/features/products"

	"gorm.io/gorm"
)

type MySqlProductRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) products.Data {
	return &MySqlProductRepository{
		Conn: conn,
	}
}

func (nr *MySqlProductRepository) SelectAllProducts(products.Products) (prd []products.Products, err error) {
	var record []Products
	err = nr.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return toProductsCoreList(record), nil
}
func (nr *MySqlProductRepository) SelectProductsById(id int) (prd products.Products, err error) {
	var idproducts Products

	err = nr.Conn.First(&idproducts, id).Error

	if idproducts.NameID == "" && idproducts.ID == 0 {
		return products.Products{}, errors.New("products not found")
	}

	if err != nil {
		return products.Products{}, err
	}

	return toProductsCore(idproducts), nil
}
func (nr *MySqlProductRepository) InsertProducts(data products.Products) (err error) {
	convData := toProductsRecord(data)

	if err := nr.Conn.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
func (nr *MySqlProductRepository) UpdateProducts(id int, data products.Products) (prd products.Products, err error) {
	var singleproducts Products
	err = nr.Conn.Model(&singleproducts).Where("id=?", id).Updates(&data).Error

	if err != nil {
		return prd, err
	}

	return toProductsCore(singleproducts), nil
}
func (nr *MySqlProductRepository) DestroyProducts(id int) (prd products.Products, err error) {
	var singleproducts Products
	fmt.Println("Isi single account : ", singleproducts)
	fmt.Println("id : ", id)
	err = nr.Conn.Model(&singleproducts).Where("id=?", id).Delete(&singleproducts).Error

	if err != nil {
		return prd, err
	}

	return toProductsCore(singleproducts), nil
}
