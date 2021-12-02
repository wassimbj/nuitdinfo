package api

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"nuitdinfo.api/config"
	"nuitdinfo.api/services"
)

type SignupData struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  string `json:"password"`
}

func Signup(c *fiber.Ctx) error {
	c.Accepts("application/json") // "application/json"

	var userInfo SignupData
	c.BodyParser(&userInfo)

	if _, exists := services.UserExists(userInfo.Email, 0); exists > 0 {
		// log.Println("USER ALREADY !!")
		c.Status(403).JSON("user already exist")
		return errors.New("user already exists")
	}

	_, err := services.CreateUser(userInfo.Firstname, userInfo.Lastname, userInfo.Email, userInfo.Password)

	if err != nil {
		// log.Println("OOOOOOOOOOOOPS")
		c.Status(500).JSON("something went wrong")
		return errors.New("something went wrong")
	}

	c.Status(200).JSON("OK !")

	return nil
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	c.Accepts("application/json") // "application/json"

	var userInfo LoginData
	c.BodyParser(userInfo)

	fmt.Println(userInfo.Email)
	existedUser, exists := services.UserExists(userInfo.Email, 0)

	if exists <= 0 {
		c.Status(403).JSON("user doesn't exist")
		return errors.New("user doesn't exist")
	}

	if !existedUser.IsAccepted {
		c.Status(403).JSON("wait until you get accepted")
		return errors.New("wait until you get accepted")
	}

	loggedInUserId, err := services.LoginUser(userInfo.Email, userInfo.Password)

	if err != nil {
		c.Status(500).JSON("something went wrong")
		return errors.New("something went wrong")
	}

	config.Session().Storage.Set("user_id", []byte(strconv.Itoa(loggedInUserId)), time.Duration(time.Now().Day()*20))
	c.Status(200).JSON("OK !")
	return nil
}
