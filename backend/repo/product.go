package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	ImgURL      string  `json:"imgURL" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(id int) error
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

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) Get(id int) (*Product, error) {
	var prd Product
	query := `SELECT * FROM products WHERE id=$1`
	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return &prd, nil
}
func (r *productRepo) List() ([]*Product, error) {
	var productList []*Product
	query := `SELECT * FROM products`
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
func (r *productRepo) Update(product Product) (*Product, error) {
	query := `
		UPDATE products
		SET 
			title = :title,
			price = :price,
			description = :description,
			img_url = :img_url
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
