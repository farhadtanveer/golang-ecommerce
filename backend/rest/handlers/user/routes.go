package user

import (
	middleware "ecommerce/rest/middleswares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// User routes
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
}