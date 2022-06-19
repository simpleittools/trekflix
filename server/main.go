package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/simpleittools/trekflix/database"
	"github.com/simpleittools/trekflix/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Conn()

	port := os.Getenv("GOSERVERPORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(port))
}
