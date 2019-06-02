package repository

import (
	"context"
	"log"

	"github.com/nantaphop/go-clean-arch/src/models"
	"github.com/nantaphop/go-clean-arch/src/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	Collection *mongo.Collection
}

// NewMongoUserRepository will create UserRepository based on Mongodb
func NewMongoUserRepository(db *mongo.Database) user.Repository {
	return &mongoUserRepository{db.Collection("users")}
}

func (r *mongoUserRepository) Create(user *models.User) (*models.User, error) {
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, nil
}
