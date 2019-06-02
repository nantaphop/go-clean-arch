package user

import "github.com/nantaphop/go-clean-arch/app/models"

// Usecase contain all operation that related to user
type Usecase interface {
	CreateUser(*models.User) (*models.User, error)
}
