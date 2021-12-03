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

	apiRoute := app.Group("api")
	apiRoute.Post("/auth/login", api.Login)
	apiRoute.Post("/rescue/add", api.CreateRescue)
	apiRoute.Post("/rescue/edit", api.EditRescue)
	apiRoute.Get("/rescue/delete/:id", api.DeleteRescue)
	apiRoute.Get("/rescue/accept/:id", api.AcceptRescue)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 👋!")
	})

	app.Listen(":8000")
}
