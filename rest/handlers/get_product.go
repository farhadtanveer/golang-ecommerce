package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request){
	productId := r.PathValue("id")

	pId, err  := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product :=database.Get(pId)
	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	util.SendData(w, product, 200)
}