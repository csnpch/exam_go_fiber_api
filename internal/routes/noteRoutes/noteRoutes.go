package noteRoutes

import (
	// noteHandler "./main/noteHandler"
	"github.com/gofiber/fiber/v2"
)

// func SetupRoutes(app *fiber.App) {
// func UserRouter (app *fiber.App) {

func SetupRoutes(router fiber.Router) {

	note := router.Group("/note")
	// Read all Notes
	// note.Get("/", noteHandler.Get())
	note.Get("/", GetNote)
	// Read one Note
	note.Get("/:noteId", func(c *fiber.Ctx) error { return nil })
	// Create a Note
	note.Post("/", func(c *fiber.Ctx) error { return nil })
	// Update one Note
	note.Put("/:noteId", func(c *fiber.Ctx) error { return nil })
	// Delete one Note
	note.Delete("/:noteId", func(c *fiber.Ctx) error { return nil })
	 
}

func GetNote(c *fiber.Ctx) error {
	return c.SendString("This note router")
}

