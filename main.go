package main

import (
	"main/database"
	"main/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	// Config
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost:4000",
	}))

	// Setup 
	database.ConnectDB()
	routes.SetupRoutes(app)

	app.Listen(":3000")

}