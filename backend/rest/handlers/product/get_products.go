package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusOK, productList) // Send the product list as JSON response
}