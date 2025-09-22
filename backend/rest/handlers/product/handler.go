package product

import middleware "ecommerce/rest/middleswares"

type Handler struct {
	middlewares *middleware.Middleswares
}

func NewHandler(middlewares *middleware.Middleswares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}