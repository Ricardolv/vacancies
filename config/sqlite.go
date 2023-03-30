package config

import (
	"os"

	"github.com/Ricardolv/vacancies/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// checked if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")

		// create database file and directory
		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migration the schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}

	// Return the database
	return db, nil
}
