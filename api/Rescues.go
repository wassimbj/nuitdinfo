package api

import (
	"github.com/gofiber/fiber/v2"
	"nuitdinfo.api/database"
)

type NewRescueData struct {
	Saver     database.Saver
	SavedUser database.SavedUser
	Boat      database.Boat
}

func CreateRescue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var rescue NewRescueData
	// var saver database.Saver
	// var savedUser database.SavedUser
	// var boat database.Boat

	// database.DB().First(&saver).Where()
	database.DB().Create(rescue)
	// databa.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)

	// database.DB().First()

	c.Status(200).JSON("ADDED !")
	return nil
}
