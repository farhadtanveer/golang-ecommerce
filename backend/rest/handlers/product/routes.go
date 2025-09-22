package product

import (
	middleware "ecommerce/rest/middleswares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Product routes
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), middleware.AuthenticateJWT))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProduct)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct), middleware.AuthenticateJWT))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct), middleware.AuthenticateJWT))
}