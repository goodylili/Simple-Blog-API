package config

import (
	"BlogAPI/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ManageMigrations() {
	DB, err = gorm.Open(sqlite.Open("Blogs.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = DB.AutoMigrate(&models.Blogs{})
	if err != nil {
		panic("migration failure")
	}
}
