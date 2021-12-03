package api

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"nuitdinfo.api/database"
)

type NewRescueData struct {
	Saver     database.Saver     `json:"saver"`
	SavedUser database.SavedUser `json:"saved_user"`
	Boat      database.Boat      `json:"boat"`
	Location  string             `json:"loc"`
	Date      string             `gorm:"type:timestamp" json:"da"`
}

func CreateRescue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var rescue NewRescueData

	c.BodyParser(&rescue)

	log.Println("DATE: ", rescue.Date)

	database.DB().Create(&rescue.Saver)
	database.DB().Create(&rescue.SavedUser)
	database.DB().Create(&rescue.Boat)

	rescueToCreate := database.Rescue{
		IdSaver:     rescue.Saver.Id,
		IdSavedUser: rescue.SavedUser.Id,
		IdBoat:      rescue.Boat.Id,
		Location:    rescue.Location,
		Date:        rescue.Date,
	}

	database.DB().Create(&rescueToCreate)

	c.Status(200).JSON("ADDED !")

	return nil
}

type EditRescueData struct {
	Saver     database.Saver     `json:"saver"`
	SavedUser database.SavedUser `json:"saved_user"`
	Boat      database.Boat      `json:"boat"`
	Location  string             `json:"location"`
	Date      time.Time          `json:"date"`
}

func EditRescue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var rescue EditRescueData

	c.BodyParser(&rescue)
	Saver := new(database.Saver)
	Saver = &rescue.Saver
	Boat := new(database.Boat)
	Boat = &rescue.Boat
	SavedUser := new(database.SavedUser)
	SavedUser = &rescue.SavedUser

	saverResult := database.DB().Model(database.Saver{}).Where(Saver.Id).Updates(Saver)
	result1 := database.DB().Model(database.Boat{}).Updates(Boat)
	result2 := database.DB().Model(database.SavedUser{}).Updates(SavedUser)
	if result1.RowsAffected < 1 || (result2.RowsAffected < 1 || saverResult.RowsAffected < 1) {
		return c.Status(202).SendString("Error no data has Updated ")
	}
	// log.Println()

	c.Status(200).JSON("DONE !")
	return nil
}

func DeleteRescue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	rescueId, _ := strconv.Atoi(c.Params("id"))

	database.DB().Delete(&database.Rescue{Id: rescueId})

	c.Status(200).JSON("DONE !")
	return nil
}

func AcceptRescue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	rescueId, _ := strconv.Atoi(c.Params("id"))

	database.DB().Update("is_accepted", 1).Where("id = ?", rescueId)

	c.Status(200).JSON("DONE !")
	return nil
}
