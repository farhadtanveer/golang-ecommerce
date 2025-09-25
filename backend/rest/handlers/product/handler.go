package product

import (
	"ecommerce/repo"
	middleware "ecommerce/rest/middleswares"
)

type Handler struct {
	middlewares *middleware.Middleswares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middleware.Middleswares, 
	productRepo repo.ProductRepo,
	) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}