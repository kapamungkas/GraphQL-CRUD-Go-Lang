package services

import (
	"graphql-product/entities"
	"graphql-product/repositories"
	"graphql-product/requests"
)

type ProductService interface {
	GetByIDProduct(id int) (interface{}, error)
	GetAllProduct() (interface{}, error)

	CreateProduct(product requests.Product) (interface{}, error)
	UpdateProduct(id int, product entities.Product) (interface{}, error)
	DeleteProduct(id int) (interface{}, error)
}

type productService struct {
	r repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) *productService {
	return &productService{r}
}

func (s productService) GetByIDProduct(id int) (interface{}, error) {
	return s.r.GetByIDProduct(id)
}

func (s productService) GetAllProduct() (interface{}, error) {
	return s.r.GetAllProduct()
}

func (s productService) CreateProduct(product entities.Product) (interface{}, error) {
	return s.r.CreateProduct(product)
}

func (s productService) UpdateProduct(id int, product entities.Product) (interface{}, error) {
	_, err := s.r.UpdateProduct(id, product)
	if err == nil {
		return s.r.GetByIDProduct(id)
	}
	return nil, err
}

func (s productService) DeleteProduct(id int) (interface{}, error) {
	return s.r.DeleteProduct(id)
}
