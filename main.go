package main

import (
	"todoapp/database"
	"todoapp/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {

	// Set up Fiber
	app := fiber.New()

	app.Use(cors.New())
	database.ConnectDB()

	router.SetupRoutes(app)
	// Start the server
	log.Fatal(app.Listen(":3000"))
}
