package service

import (
	"MyEcommerce/features/product"
	"errors"
)

type productService struct {
	productData product.ProductDataInterface
}

func New(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: repo,
	}
}

func (ps *productService) Create(userIdLogin int, input product.Core) error {
	if input.Name == "" {
		return errors.New("nama produk tidak boleh kosong")
	}

	if input.Price <= 0 {
		return errors.New("harga produk harus lebih besar dari 0")
	}

	err := ps.productData.Insert(userIdLogin, input)
	if err != nil {
		return err
	}

	return nil
}

// GettAll implements product.ProductServiceInterface.
func (ps *productService) GetAll(page, limit int, category string) ([]product.Core, int, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 8
	}

	products, totalPage, err := ps.productData.SelectAll(page, limit, category)
	if err != nil {
		return nil, 0, err
	}

	return products, totalPage, nil
}

// GetById implements product.ProductServiceInterface.
func (ps *productService) GetById(IdProduct int) (*product.Core, error) {
	result, err := ps.productData.SelectById(IdProduct)
	return result, err
}

// Update implements product.ProductServiceInterface.
func (ps *productService) Update(userIdLogin int, input product.Core) error {
	err := ps.productData.Update(userIdLogin, input)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements product.ProductServiceInterface.
func (ps *productService) Delete(userIdLogin int, IdProduct int) error {
	err := ps.productData.Delete(userIdLogin, IdProduct)
	if err != nil {
		return err
	}

	return nil
}

// GetByUserId implements product.ProductServiceInterface.
func (ps *productService) GetByUserId(userIdLogin int) ([]product.Core, error) {
	products, err := ps.productData.SelectByUserId(userIdLogin)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Search implements product.ProductServiceInterface.
func (ps *productService) Search(query string) ([]product.Core, error) {
	products, err := ps.productData.Search(query)
	if err != nil {
		return products, err
	}
	return products, nil
}
