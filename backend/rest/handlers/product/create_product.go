package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgURL      string  `json:"imgURL"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	createdProduct, err := h.svc.Create(domain.Product{
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		ImgURL:      req.ImgURL,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Error creating product")
		return
	}

	util.SendData(w, http.StatusCreated,createdProduct) 
}