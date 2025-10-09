package product

import (
	middleware "ecommerce/rest/middleswares"
)

type Handler struct {
	middlewares *middleware.Middleswares
	svc 	  Service
}

func NewHandler(
	middlewares *middleware.Middleswares, 
	svc Service,
	) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc: 		 svc,
	}
}