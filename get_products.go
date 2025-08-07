package main

import (
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, productList, 200) // Send the product list as JSON response
}