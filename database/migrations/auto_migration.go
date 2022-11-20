package migrations

import (
	"depoguna-api/models"

	"gorm.io/gorm"
)

// AutoMigration is auto migrate database
func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Customer{}, &models.Order{})
}
