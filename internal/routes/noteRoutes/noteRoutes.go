package noteRoutes

import (
	"github.com/gofiber/fiber/v2"

	noteHandler "main/internal/handlers/noteHandler"
)

// func SetupRoutes(app *fiber.App) {
// func UserRouter (app *fiber.App) {

func SetupRoutes(router fiber.Router) {

	note := router.Group("/note")
	// note.Get("/", noteHandler.Get())
	note.Get("/", noteHandler.GetNote)
	// Read one Note
	note.Get("/:noteId", noteHandler.GetNote)
	// Create a Note
	note.Post("/", noteHandler.CreateNotes)
	// Update one Note
	note.Put("/:noteId", noteHandler.UpdateNote)
	// Delete one Note
	note.Delete("/:noteId", noteHandler.DeleteNote)
	 
}
