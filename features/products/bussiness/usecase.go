package bussiness

import "java-agro/features/products"

type ProductUsecase struct {
	pData products.Data
}

func NewBussinessProduct(prdData products.Data) products.Bussiness {
	return &ProductUsecase{
		pData: prdData,
	}
}

func (pu *ProductUsecase) GetAllProducts(data products.Products) (prd []products.Products, err error) {
	products, err := pu.pData.SelectAllProducts(data)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (pu *ProductUsecase) GetProductsById(id int) (prd products.Products, err error) {
	productsData, err := pu.pData.SelectProductsById(id)

	if err != nil {
		return products.Products{}, err
	}
	return productsData, nil
}
func (pu *ProductUsecase) CreateProducts(data products.Products) (err error) {
	if err := pu.pData.InsertProducts(data); err != nil {
		return err
	}
	return nil
}
func (pu *ProductUsecase) EditProducts(id int, data products.Products) (prd products.Products, err error) {
	productsData, err := pu.pData.UpdateProducts(id, data)

	if err != nil {
		return prd, err
	}
	return productsData, nil
}
func (pu *ProductUsecase) DeleteProducts(id int) (prd products.Products, err error) {
	productsData, err := pu.pData.DestroyProducts(id)

	if err != nil {
		return prd, err
	}
	return productsData, nil
}
