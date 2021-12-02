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

	//? create account
	// app.Post("/auth/signup", api.Signup)
	app.Post("/auth/login", api.Login)
	app.Post("/rescue/add", api.CreateRescue)
	// app.Post("/auth/login", api.Login)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 👋!")
	})

	app.Listen(":8000")
}
