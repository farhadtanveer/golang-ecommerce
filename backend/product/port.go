package product

import (
	"ecommerce/domain"
	prdctHandler "ecommerce/rest/handlers/product"
)

type Service interface {
	prdctHandler.Service
}

type ProductRepo interface {
	Create(domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Update(domain.Product) (*domain.Product, error)
	Delete(id int) error
} 