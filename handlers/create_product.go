package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	newProduct.ID = len(database.ProductList) + 1 // Assign a new ID based on the current length of the product list
	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 200) 
}