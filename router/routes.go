package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	noteRoutes "main/internal/routes/noteRoutes"
	userRoutes "main/internal/routes/userRoutes"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	userRoutes.SetupRoutes(api)
	noteRoutes.SetupRoutes(api)

}


