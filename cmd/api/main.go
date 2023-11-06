package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spiandorello/todo-list-golang/src/handlers"
	"github.com/spiandorello/todo-list-golang/src/repositories"
	"github.com/spiandorello/todo-list-golang/src/services"
	"github.com/spiandorello/todo-list-golang/src/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=todo password=todo dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&structs.User{}, &structs.Task{})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	repositoryContainer := repositories.GetContainer(db)
	servicesContainer := services.GetContainer(repositoryContainer)

	handlers.New(servicesContainer).
		GetRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
