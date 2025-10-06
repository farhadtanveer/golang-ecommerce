package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)


type Service interface {
	userHandler.Service // embedding
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	// Update(user User) (*User, error)
	// Delete(userID int) error
}