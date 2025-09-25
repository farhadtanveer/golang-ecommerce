package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgURL      string  `json:"imgURL"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request){
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:          pId,
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		ImgURL:      req.ImgURL,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Error updating product")
		return
	}
	util.SendData(w, http.StatusOK,"Successfully updated product")
}