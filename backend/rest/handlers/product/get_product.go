package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request){
	productId := r.PathValue("id")

	pId, err  := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err :=  h.svc.Get(pId)
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}
	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	util.SendData(w, http.StatusOK, product)
}