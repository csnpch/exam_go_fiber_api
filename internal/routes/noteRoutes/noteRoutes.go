package noteRoutes

import (
	"github.com/gofiber/fiber/v2"

	noteController "main/internal/controllers"
)

// func SetupRoutes(app *fiber.App) {
// func UserRouter (app *fiber.App) {

func SetupRoutes(router fiber.Router) {
	
	// v1 := router.Group("v1")
	note := router.Group("/note")
	// Get
	note.Get("/", noteController.GetNote)
	// Read one Note
	note.Get("/:noteId", noteController.GetNote)
	// Create a Note
	note.Post("/", noteController.CreateNotes)
	 
}
