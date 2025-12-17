package main

import (
	"log"

	"github.com/SubhaNITS150/dob-calculator/config"
	"github.com/SubhaNITS150/dob-calculator/internal/handler"
	"github.com/SubhaNITS150/dob-calculator/internal/repository"
	"github.com/SubhaNITS150/dob-calculator/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.Load()

	db, err := config.NewDB(cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully")

	repo := repository.NewUserRepository(db)
	handler := handler.NewUserHandler(repo)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	routes.Setup(app, handler)

	log.Println("Server listening on port 8090")
	log.Fatal(app.Listen(":8090"))
}
