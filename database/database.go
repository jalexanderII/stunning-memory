package database

import (
	"github.com/jalexanderII/stunning-memory/config"
	"github.com/jalexanderII/stunning-memory/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DbInstance is a struct that holds database pointer
type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dbLogger := config.Logger.Named("database")

	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		dbLogger.Error("failed to connect database", "error", err)
	}
	dbLogger.Info("Connecting to database")

	// Migrate the schema
	dbLogger.Info("Running Migrations")
	db.AutoMigrate(&models.Order{}, &models.Order{}, &models.Product{})

	Database = DbInstance{Db: db}
}