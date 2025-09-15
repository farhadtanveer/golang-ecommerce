package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200) // Send the product list as JSON response
}