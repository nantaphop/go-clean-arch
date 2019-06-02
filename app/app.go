package app

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nantaphop/go-clean-arch/app/core"
	userRestDelivery "github.com/nantaphop/go-clean-arch/app/user/delivery/rest"
	userRepository "github.com/nantaphop/go-clean-arch/app/user/repository"
	userUsecase "github.com/nantaphop/go-clean-arch/app/user/usecase"
)

// InitApp will initial application
func InitApp() {
	client, err := core.GetMongo()
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	log.Println("Connected to MongoDB!")
	// Prepared Mongo Connection
	defer client.Disconnect(context.Background())
	db := client.Database("go-clean-arch")

	// Setup dependencies
	userRepo := userRepository.NewMongoUserRepository(db)
	userUc := userUsecase.NewUserUsecase(userRepo)

	// Setup Rest Delivery
	r := gin.Default()
	userRestDelivery.NewUserDelivery(r, userUc)
	r.Run(":3000")
}
