package services

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"nuitdinfo.api/database"
)

func UserExists(email string, id int) (database.Admin, int64) {
	// by email
	var user database.Admin
	var affectedRows int64
	if email != "" && id == 0 {
		result := database.DB().Select("id", "password").Where(&database.Admin{
			Email: email,
		}).First(&user)
		affectedRows = result.RowsAffected
	} else {
		result := database.DB().Select("id").Where(&database.Admin{
			Id: id,
		}).First(&user)
		affectedRows = result.RowsAffected
	}

	return user, affectedRows
}

// returns user id and error
func CreateUser(first_name string, last_name string, email string, password string) (int, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	newUser := database.Admin{
		Firstname: first_name,
		Lastname:  last_name,
		Email:     email,
		Password:  string(hashedPass),
	}

	result := database.DB().Create(&newUser)

	if result.Error != nil {
		return 0, result.Error
	}

	return newUser.Id, nil
}

func LoginUser(email string, password string) (int, error) {

	user, affectedRows := UserExists(email, 0)
	if affectedRows == 0 {
		// record not found
		return 0, errors.New("User email not found")
	}
	matchFailed := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if matchFailed != nil {
		return 0, errors.New("Password doesn't match")
	}

	return user.Id, nil
}

func GetLoggedInUser(userId interface{}) (database.Admin, bool) {
	var user database.Admin

	result := database.DB().Select("id", "firstname").Where(&database.Admin{Id: userId.(int)}).First(&user)
	if result.RowsAffected == 0 {
		return user, false
	}
	return user, true
}
