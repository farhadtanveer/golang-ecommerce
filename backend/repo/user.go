package repo

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
    ID          int       `db:"id" json:"id"`
    FirstName   string    `db:"first_name" json:"first_name"`
    LastName    string    `db:"last_name" json:"last_name"`
    Email       string    `db:"email" json:"email"`
    Password    string    `db:"password" json:"password"`
    IsShopOwner bool      `db:"is_shop_owner" json:"is_shop_owner"`
}


type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
	// List() ([]*User, error)
	// Update(user User) (*User, error)
	// Delete(userID int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	 query := `
        INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			is_shop_owner
		)
        VALUES (
			:first_name, 
			:last_name, 
			:email, 
			:password, 
			:is_shop_owner
		)
        RETURNING id
    `
	// execute named query
	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&userID)
	}
	
	user.ID = userID
    return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	query := `
		SELECT 
			id, 
			first_name, 
			last_name, 
			email, 
			password, 
			is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
	`
	var user User
	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}