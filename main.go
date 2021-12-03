package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"nuitdinfo.api/api"
)

func main() {
	app := fiber.New()

	//? setup cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	apiRoute := app.Group("api")
	apiRoute.Post("/auth/login", api.Login)
	apiRoute.Get("/auth/isauth", api.GetLoggedInAdmin)
	apiRoute.Post("/rescue/add", api.CreateRescue)
	apiRoute.Post("/rescue/edit", api.EditRescue)
	apiRoute.Delete("/rescue/delete/:id", api.DeleteRescue)
	apiRoute.Put("/rescue/accept/:id", api.AcceptRescue)
	apiRoute.Get("/rescue/saver/:name", api.GetSaver)
	apiRoute.Get("/rescue/saved/:name", api.GetSaveds)
	apiRoute.Get("/rescue", api.GetRescue)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 👋!")
	})

	app.Listen(":8000")
}
