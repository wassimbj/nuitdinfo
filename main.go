package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New() //! app = express()

	app.Get("/", func(c *fiber.Ctx) error { //! app.get(...)
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
