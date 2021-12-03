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
	Firstname   string `gorm:"size:50" json:"firstname"`
	Lastname    string `gorm:"size:50" json:"lastname"`
	Email       string `gorm:"size:255" json:"email"`
	Password    string `gorm:"size:80" json:"password"`
	VerifyToken string `gorm:"size:50" json:"verifytoken"`
	IsAccepted  bool   `gorm:"size:50" json:"isaccepted"` // is accepted by the admin ?
	IsVerified  bool   `gorm:"size:50" json:"isverified"` // is email verified
	DN          string `gorm:"type:date" json:"dn"`
	Role        string `sql:"type:ENUM('simple_user', 'savior')"`
}

// simple user
type Saver struct {
	Id        int    `gorm:"autoIncrement"`
	Firstname string `gorm:"size:50" json:"first_name"`
	Lastname  string `gorm:"size:50" json:"last_name"`
	Age       int    `gorm:"size:50" json:"age"`
}

type SavedUser struct {
	Id        int    `gorm:"autoIncrement"`
	Firstname string `gorm:"size:50" json:"first_name"`
	Lastname  string `gorm:"size:50" json:"last_name"`
	Age       int    `gorm:"size:50" json:"age"`
	Length    int    `json:"length"`
	Taille    int    `json:"taille"`
	State     int    `json:"state"`
}

type Boat struct {
	Id    int    `gorm:"autoIncrement"`
	Name  string `gorm:"size:50" json:"name"`
	State int    `json:"state"`
}

type Rescue struct {
	Id          int `gorm:"autoIncrement"`
	IdSaver     int `gorm:"foreignKey:"`
	IdSavedUser int `gorm:"foreignKey:"`
	IdBoat      int
	Date        time.Time `gorm:"type:timestamp" json:"data"`
	IsAccepted  bool
	Location    string
}
type ReqRescue struct {
	Saver    Saver
	Persons  []SavedUser
	Boat     Boat
	Date     time.Time
	Location string
}
