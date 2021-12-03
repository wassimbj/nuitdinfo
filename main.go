package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"nuitdinfo.api/api"
)

func main() {
	app := fiber.New()

	//? setup cors
	app.Use(cors.New())

	app.Post("/auth/login", api.Login)
	app.Post("/rescue/add", api.CreateRescue)
	app.Post("/rescue/edit", api.EditRescue)
	app.Get("/rescue/delete/:id", api.DeleteRescue)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 👋!")
	})

	app.Listen(":8000")
}
