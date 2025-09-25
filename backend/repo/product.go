package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgURL      string  `json:"imgURL"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(id int) error
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
// ekta function jeita struct er ekta instance create kore dey
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(id int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(id int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != id {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Product 1",
		Price:       19.99,
		Description: "This is the first product",
		ImgURL:      "http://example.com/product1.jpg",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Product 2",
		Price:       29.99,
		Description: "This is the second product",
		ImgURL:      "http://example.com/product1.jpg",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
}