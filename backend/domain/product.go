package domain

// model or entity -> existence
type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	ImgURL      string  `json:"imgURL" db:"img_url"`
}