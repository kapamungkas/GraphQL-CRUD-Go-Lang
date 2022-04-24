package repositories

import (
	"graphql-product/entities"
	"graphql-product/requests"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() (interface{}, error)
	GetByIDProduct(id int) (interface{}, error)
	CreateProduct(product entities.Product) (interface{}, error)
	UpdateProduct(id int, product entities.Product) (interface{}, error)
	DeleteProduct(id int) (interface{}, error)
}

type productRepository struct {
	databBase *gorm.DB
}

var Products = []requests.Product{
	{
		ID:    1,
		Name:  "Chicha Morada",
		Info:  "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)",
		Price: 7.99,
	},
	{
		ID:    2,
		Name:  "Chicha de jora",
		Info:  "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)",
		Price: 5.95,
	},
	{
		ID:    3,
		Name:  "Pisco",
		Info:  "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)",
		Price: 9.95,
	},
}

func NewProductRepository(databBase *gorm.DB) *productRepository {
	return &productRepository{databBase}
}

func (t productRepository) GetByIDProduct(id int) (interface{}, error) {
	// Find product
	// fmt.Println(Products)
	// for _, product := range Products {
	// 	if int(product.ID) == id {
	// 		return product, nil
	// 	}
	// }
	// return nil, nil

	var product entities.Product

	err := t.databBase.Find(&product, id).Error

	return product, err
}

func (t productRepository) GetAllProduct() (interface{}, error) {
	var products []entities.Product

	err := t.databBase.Find(&products).Error

	return products, err

	// fmt.Println(Products)
	// return Products, nil
}

func (t productRepository) CreateProduct(product entities.Product) (interface{}, error) {
	err := t.databBase.Create(&product).Error
	return product, err
}

func (t productRepository) UpdateProduct(id int, product entities.Product) (interface{}, error) {
	var products entities.Product

	err := t.databBase.Model(&products).Where("id = ?", id).Updates(product).Error

	return product, err
}

func (t productRepository) DeleteProduct(id int) (interface{}, error) {
	var products entities.Product

	err := t.databBase.Where("id = ?", id).Delete(&products).Error

	return products, err
}
