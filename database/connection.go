package database

import (
	"data_app/api"
	"data_app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

var modelList = []any{
	models.Customer{},
}

func Connect() error {
	var err error

	DB, err = gorm.Open(sqlite.Open(api.Config.Database.DatabaseName), &gorm.Config{
		PrepareStmt: api.Config.Database.PrepareStmt,
	})
	if err != nil {
		return err
	}

	DB.AutoMigrate(modelList...)

	return nil
}
