package main

import (
	"context"
	"log"

	"github.com/nantaphop/go-clean-arch/src/core"
	"github.com/nantaphop/go-clean-arch/src/models"
	"github.com/nantaphop/go-clean-arch/src/user/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	client, err := core.GetMongo()
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	log.Println("Connected to MongoDB!")
	defer client.Disconnect(context.Background())
	db := client.Database("go-clean-arch")
	userRepo := repository.NewMongoUserRepository(db)
	result, err := userRepo.Create(&models.User{
		Username: "Test",
		Password: "sddf",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created %s", result)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "vv",
		})
	})
	r.Run(":3000")
}
