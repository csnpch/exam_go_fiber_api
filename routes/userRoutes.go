package routes

import (
	"github.com/gofiber/fiber/v2"
)


func UserRouter (app *fiber.App) { 
	
	api := app.Group("/api")
	v1 := api.Group("/v1")
	
	v1.Get("/users", GetAllUsers)

}


func GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("GetAllUser working...")
}