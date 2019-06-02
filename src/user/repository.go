package user

import (
	"github.com/nantaphop/go-clean-arch/src/models"
)

// Repository handle basic operation such as CRUD
type Repository interface {
	Create(*models.User) (*models.User, error)
}
