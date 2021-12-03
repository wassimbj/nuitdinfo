package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"nuitdinfo.api/database"
)

type NewRescueData struct {
	Saver     database.Saver       `json:"saver"`
	SavedUser []database.SavedUser `json:"saved_user"`
	Boat      database.Boat        `json:"boat"`
	Location  string               `json:"loc"`
	Date      string               `gorm:"size:20" json:"date"`
}

func CreateRescue(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var rescue NewRescueData

	c.BodyParser(&rescue)

	database.DB().Create(&rescue.Saver)
	database.DB().Create(&rescue.SavedUser)
	database.DB().Create(&rescue.Boat)

	fmt.Println("DATE: ", rescue.Date)
	for _, i := range rescue.SavedUser {
		database.DB().Create(&database.Rescue{
			IdSaver:     rescue.Saver.Id,
			IdSavedUser: i.Id,
			IdBoat:      rescue.Boat.Id,
			Location:    rescue.Location,
			Date:        rescue.Date,
		})
	}

	c.Status(200).JSON("ADDED !")

	return nil
}

type EditRescueData struct {
	Saver     database.Saver     `json:"saver"`
	SavedUser database.SavedUser `json:"saved_user"`
	Boat      database.Boat      `json:"boat"`
	Location  string             `json:"location"`
	Date      string             `json:"date"`
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

	// database.DB().Update("is_accepted", 1).Where("id = ?", rescueId)
	database.DB().Model(&database.Rescue{}).Where("id = ?", rescueId).Update("is_accepted", 1)

	c.Status(200).JSON("DONE !")
	return nil
}
func GetSaver(c *fiber.Ctx) error {

	saverName := c.Params("name")
	saver := new(database.Saver)
	saver.Firstname = saverName
	database.DB().Where("firstname", saver.Firstname).Find(&saver)
	if saver.Id < 1 {
		return c.Status(fiber.StatusNotFound).JSON("Data not found")
	}
	return c.Status(fiber.StatusOK).JSON(saver)
}
func GetSaveds(c *fiber.Ctx) error {

	savedName := c.Params("name")
	Saved := new(database.SavedUser)
	Saved.Firstname = savedName
	database.DB().Where("firstname", Saved.Firstname).Find(&Saved)
	if Saved.Id < 1 {
		return c.Status(fiber.StatusNotFound).JSON("Data not found")
	}
	return c.Status(fiber.StatusOK).JSON(Saved)
}
func GetRescue(c *fiber.Ctx) error {
	rescue := new([]database.Rescue)
	database.DB().Find(&rescue)
	if len(*rescue) < 1 {
		return c.Status(fiber.StatusNotFound).JSON("Data not found")
	}
	return c.Status(fiber.StatusOK).JSON(rescue)
}
