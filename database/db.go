package db

import (

	// "gorm.io/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func init() {
	DSN := "root:12345@tcp(nuitdinfo_db:3306)/nuitdinfo"

	var pg = postgres.Open(DSN)
	var err error
	dbConn, err = gorm.Open(pg, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		recover()
		panic("failed to connect database")
	}
	// migrate database schema
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Link{})
}

func DB() *gorm.DB {
	return dbConn.Debug()
}
