package userRoutes

import (
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(router fiber.Router) {
	// v1 := router.Group("v1")
	router.Get("/users", GetAllUsers)
}


func GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("GetAllUser working...")
}
