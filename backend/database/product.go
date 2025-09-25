package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgURL      string  `json:"imgURL"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(id int) *Product {
	for _, product := range productList {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
			return
		}
	}
}

func Delete(id int) {
	var tempList []Product = make([]Product, 0)

	for _, p := range productList {
		if p.ID != id {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func generateInitialProducts() {
	prd1 := Product{
		ID:          1,
		Title:       "Product 1",
		Price:       19.99,
		Description: "This is the first product",
		ImgURL:      "http://example.com/product1.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Product 2",
		Price:       29.99,
		Description: "This is the second product",
		ImgURL:      "http://example.com/product1.jpg",
	}

	productList = append(productList, prd1, prd2)
}