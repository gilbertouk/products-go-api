package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = id
	return product, nil
}

func (pu *ProductUseCase) GetProductByID(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(product model.Product) error {
	_, err := pu.repository.GetProductByID(product.ID)
	if err != nil {
		return err
	}
	err = pu.repository.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (pu *ProductUseCase) DeleteProduct(id int) error {
	return pu.repository.DeleteProduct(id)
}
