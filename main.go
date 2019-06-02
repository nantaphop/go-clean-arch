package main

import "github.com/nantaphop/go-clean-arch/app"

func main() {
	// client, err := core.GetMongo()
	// if err != nil {
	// 	log.Fatal("Can't connect to database")
	// }
	// log.Println("Connected to MongoDB!")
	// defer client.Disconnect(context.Background())
	// db := client.Database("go-clean-arch")
	// userRepo := userRepository.NewMongoUserRepository(db)
	// userUc := userUsecase.NewUserUsecase(userRepo)

	// result, err := userUc.CreateUser(&models.User{
	// 	Username: "Test",
	// 	Password: "sddf",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Created %s", result)
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "vv",
	// 	})
	// })
	// r.Run(":3000")

	app.InitApp()
}
