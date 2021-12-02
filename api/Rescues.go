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

	database.DB().Create(rescue.Saver)
	database.DB().Create(rescue.SavedUser)
	database.DB().Create(rescue.Boat)

	database.DB().Create(database.Rescue{
		IdSaver:     rescue.Saver.Id,
		IdSavedUser: rescue.SavedUser.Id,
		IdBoat:      rescue.Boat.Id,
	})

	c.Status(200).JSON("ADDED !")
	return nil
}
