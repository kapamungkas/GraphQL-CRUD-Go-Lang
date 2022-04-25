package repositories

import (
	"graphql-product/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetByIDProduct(id int) (entities.Product, error)
	GetAllProduct() ([]entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(id int, product entities.Product) (entities.Product, error)
	DeleteProduct(id int) (entities.Product, error)
}

type productRepository struct {
	databBase *gorm.DB
}

func NewProductRepository(databBase *gorm.DB) *productRepository {
	return &productRepository{databBase}
}

func (t productRepository) GetByIDProduct(id int) (entities.Product, error) {

	var product entities.Product

	err := t.databBase.Find(&product, id).Error

	return product, err
}

func (t productRepository) GetAllProduct() ([]entities.Product, error) {
	var products []entities.Product

	err := t.databBase.Find(&products).Error

	return products, err
}

func (t productRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	err := t.databBase.Create(&product).Error
	return product, err
}

func (t productRepository) UpdateProduct(id int, product entities.Product) (entities.Product, error) {
	var products entities.Product

	err := t.databBase.Model(&products).Where("id = ?", id).Updates(product).Error

	return product, err
}

func (t productRepository) DeleteProduct(id int) (entities.Product, error) {
	var products entities.Product

	err := t.databBase.Where("id = ?", id).Delete(&products).Error

	return products, err
}
