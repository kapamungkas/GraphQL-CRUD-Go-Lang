package services

import (
	"graphql-product/entities"
	"graphql-product/repositories"
)

type ProductService interface {
	GetByIDProduct(id int) (entities.Product, error)
	GetAllProduct() ([]entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(id int, product entities.Product) (entities.Product, error)
	DeleteProduct(id int) (entities.Product, error)
}

type productService struct {
	r repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) *productService {
	return &productService{r}
}

func (s productService) GetByIDProduct(id int) (entities.Product, error) {
	return s.r.GetByIDProduct(id)
}

func (s productService) GetAllProduct() ([]entities.Product, error) {
	return s.r.GetAllProduct()
}

func (s productService) CreateProduct(product entities.Product) (entities.Product, error) {
	return s.r.CreateProduct(product)
}

func (s productService) UpdateProduct(id int, product entities.Product) (entities.Product, error) {
	result, err := s.r.UpdateProduct(id, product)
	if err == nil {
		return s.r.GetByIDProduct(id)
	}
	return result, err
}

func (s productService) DeleteProduct(id int) (entities.Product, error) {
	return s.r.DeleteProduct(id)
}
