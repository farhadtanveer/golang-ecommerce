package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newProduct.ID = pId
	database.Update(newProduct)
	util.SendData(w, "Successfully updated product", 201)
}