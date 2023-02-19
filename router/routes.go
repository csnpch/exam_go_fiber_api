package routes

import (
	"github.com/gofiber/fiber/v2"

	noteRoutes "main/internal/routes/noteRoutes"
	userRoutes "main/internal/routes/userRoutes"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	userRoutes.SetupRoutes(api)
	noteRoutes.SetupRoutes(api)

}


