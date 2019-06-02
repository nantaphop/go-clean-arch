package usecase

import (
	"errors"

	"github.com/nantaphop/go-clean-arch/app/models"
	"github.com/nantaphop/go-clean-arch/app/user"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser will validate and encrypt password then save to repository

type userUsecase struct {
	repo user.Repository
}

// NewUserUsecase will return object of UserUsecase
func NewUserUsecase(r user.Repository) user.Usecase {
	return &userUsecase{r}
}

func (uc userUsecase) CreateUser(u *models.User) (*models.User, error) {
	if u.Username == "" {
		return nil, errors.New("Username is required")
	}
	if u.Password == "" {
		return nil, errors.New("Password is required")
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(newPassword)
	if err != nil {
		return nil, errors.New("Error on encrypt password")
	}
	_, err = uc.repo.Create(u)
	if err != nil {
		return nil, err
	}
	return u, nil

}
