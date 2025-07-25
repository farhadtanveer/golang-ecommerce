package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "I am Tanveer, a software developer with a passion for building web applications.")
}

type Product struct {
	ID    int          `json:"id"`
	Title  string      `json:"title"`
	Price float64      `json:"price"`
	Description string `json:"description"`
	ImgURL string 	   `json:"imgURL"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "GET" {
		http.Error(w, "Please give me get request", 400)
		return
	}

	encoder :=json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	// This is the entry point of the application.
	// You can initialize your application here.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", mux) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init(){
	prd1 := Product{
		ID:          1,
		Title:       "Product 1",
		Price:      19.99,
		Description: "This is the first product",
		ImgURL:      "http://example.com/product1.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Product 2",
		Price:      29.99,
		Description: "This is the second product",
		ImgURL:      "http://example.com/product1.jpg",
	}

	productList = append(productList, prd1, prd2)
}