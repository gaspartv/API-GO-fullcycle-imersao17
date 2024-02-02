package service

import (
	"github.com/gaspartv/API-GO-fullcycle-imersao17/internal/database"
	"github.com/gaspartv/API-GO-fullcycle-imersao17/internal/entity"
)

type ProductService struct {
	ProductDb database.ProductDB
}

func NewProductService(productDb database.ProductDB) *ProductService {
	return &ProductService{ProductDb: productDb}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDb.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDb.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDb.GetProductByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := ps.ProductDb.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
