package repo

import (
	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)


type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
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

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
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
	var user domain.User
	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}