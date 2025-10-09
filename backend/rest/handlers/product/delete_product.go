package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.svc.Delete(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Error deleting product")
		return
	}
	util.SendData(w, http.StatusOK, "Successfully deleted product")
}