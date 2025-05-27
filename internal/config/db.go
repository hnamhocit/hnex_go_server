package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hnex.com/internal/models"
)

func ConnectDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
