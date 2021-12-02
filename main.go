package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"nuitdinfo.api/config"
)

func main() {
	app := fiber.New() //? app = express()

	//? setup cors
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println(config.GetEnv("MYSQL_ROOT_PASSWORD"))
		return c.SendString("Hello 👋!")
	})

	app.Listen(":8000")
}
