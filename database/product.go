package database

var ProductList []Product

type Product struct {
	ID    int          `json:"id"`
	Title  string      `json:"title"`
	Price float64      `json:"price"`
	Description string `json:"description"`
	ImgURL string 	   `json:"imgURL"`
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

	ProductList = append(ProductList, prd1, prd2)
}