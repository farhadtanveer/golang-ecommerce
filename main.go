package main

import (
	"fmt"
	"net/http"
)

type Product struct {
	ID    int          `json:"id"`
	Title  string      `json:"title"`
	Price float64      `json:"price"`
	Description string `json:"description"`
	ImgURL string 	   `json:"imgURL"`
}

var productList []Product


func handleCors(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-Type", "application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		handleCors(w) // Set CORS headers
		w.WriteHeader(200)
		return
	}
}

func main() {
	// This is the entry point of the application.
	// You can initialize your application here.
	mux := http.NewServeMux()
	mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProducts)))
	mux.HandleFunc("POST /create-product", http.HandlerFunc(createProduct))

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

func corsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("content-Type", "application/json")
		
		next.ServeHTTP(w, r) // Call the next handler in the chain
	}
	handler := http.HandlerFunc(handleCors)
	return handler
}