package routes

import (
	"github.com/gofiber/fiber/v2"
)


func ProductRouter (app *fiber.App) { 

	api := app.Group("/api")
	v1 := api.Group("/v1")
	
	v1.Get("/products", GetAllProducts)

}


func GetAllProducts(c *fiber.Ctx) error {
	return c.SendString("GetAllProduct working...");
}

