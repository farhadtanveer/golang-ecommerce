package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)            // Set CORS headers
	handlePreflightReq(w, r) // Handle preflight request

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	newProduct.ID = len(productList) + 1 // Assign a new ID based on the current length of the product list
	productList = append(productList, newProduct)

	sendData(w, newProduct, 200) 
}