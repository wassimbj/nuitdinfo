package database

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
	Role        string `sql:"type:ENUM('production', 'testsystem')"`
}

// simple user
// type Saviors struct {
// 	Id        int    `gorm:"autoIncrement"`
// 	Firstname string `gorm:"size:50"`
// 	Lastname  string `gorm:"size:50"`
// 	Email     string `gorm:"size:255"`
// 	Password  string `gorm:"size:80"`
// }
