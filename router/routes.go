package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	userRoutes "main/internal/routes/userRoutes"
	noteRoutes "main/internal/routes/noteRoutes"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	userRoutes.SetupRoutes(api)
	noteRoutes.SetupRoutes(app)

}


