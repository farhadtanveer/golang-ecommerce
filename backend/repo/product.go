package repo

import (
	"database/sql"

	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)


type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
// ekta function jeita struct er ekta instance create kore dey
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (
			title,
			price,
			description,
			img_url
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(query, p.Title, p.Price, p.Description, p.ImgURL).Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product
	query := `SELECT * FROM products WHERE id=$1`
	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return &prd, nil
}
func (r *productRepo) List() ([]*domain.Product, error) {
	var productList []*domain.Product
	query := `
SELECT id, title, price, description, img_url
FROM products
`
	err := r.db.Select(&productList, query)
	if err != nil {
		return nil, err
	}
	return productList, nil
}
func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET 
			title = $1,
			price = $2,
			description = $3,
			img_url = $4
		WHERE id = $5
	`
	_, err := r.db.NamedExec(query, product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
