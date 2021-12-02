package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func init() {
	DSN := "root:12345@tcp(nuitdinfo_db:3306)/nuitdinfo"

	var err error

	dbConn, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("DB FAILED TO CONNECT !!")
		// recover()
		// panic("failed to connect database")
	}
	// migrate database schema
	// dbConn.AutoMigrate(&User{})
	dbConn.AutoMigrate(&Admin{})
	dbConn.AutoMigrate(&Saver{})
	dbConn.AutoMigrate(&SavedUser{})
	dbConn.AutoMigrate(&Boat{})
	dbConn.AutoMigrate(&Sauvage{})
}

func DB() *gorm.DB {
	return dbConn.Debug()
}
