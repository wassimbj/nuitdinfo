package database

import "time"

type Admin struct {
	Id        int    `gorm:"autoIncrement"`
	Firstname string `gorm:"size:50"`
	Lastname  string `gorm:"size:50"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:80"`
}

// simple user
type User struct {
	Id          int    `gorm:"autoIncrement"`
	Firstname   string `gorm:"size:50"`
	Lastname    string `gorm:"size:50"`
	Email       string `gorm:"size:255"`
	Password    string `gorm:"size:80"`
	VerifyToken string `gorm:"size:50"`
	IsAccepted  bool   `gorm:"size:50"` // is accepted by the admin ?
	IsVerified  bool   `gorm:"size:50"` // is email verified
	DN          string `gorm:"type:date"`
	Role        string `sql:"type:ENUM('simple_user', 'savior')"`
}

// simple user
type Saver struct {
	Id        int    `gorm:"autoIncrement"`
	Firstname string `gorm:"size:50"`
	Lastname  string `gorm:"size:50"`
	Age       int
}

type SavedUser struct {
	Id        int    `gorm:"autoIncrement"`
	Firstname string `gorm:"size:50"`
	Lastname  string `gorm:"size:50"`
	Age       int
	Length    int
	Taille    int
	State     int
}

type Boat struct {
	Id    int    `gorm:"autoIncrement"`
	Name  string `gorm:"size:50"`
	State int
}

type Sauvage struct {
	Id          int `gorm:"autoIncrement"`
	IdSaver     int
	IdSavedUser int
	Date        time.Time
	Location    string
}

// type SavedUser struct {
// 	Firstname string `gorm:"size:50"`
// 	Lastname  string `gorm:"size:50"`
// }

// type Boat struct {
// 	Id        int `gorm:"autoIncrement"`
// 	Matricule string
// 	TripFile  string
// 	Name      string
// }

// type TripFiles struct {
// 	Id     int `gorm:"autoIncrement"`
// 	BoatId int
// }

// type Saved struct {
// }

// type Boat struct {
// 	Id        int `gorm:"autoIncrement"`
// 	Matricule string
// 	TripFile  string
// 	Name      string
// }

// type TripFiles struct {
// 	Id     int `gorm:"autoIncrement"`
// 	BoatId int
// }

// type Saved struct {
// }
