package user

import (
	"github.com/nantaphop/go-clean-arch/app/models"
)

// Repository handle basic operation such as CRUD
type Repository interface {
	Create(*models.User) (*models.User, error)
}
