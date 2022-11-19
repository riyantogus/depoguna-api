package main

import (
	"depoguna-api/config"
	"depoguna-api/database/migrations"
	"depoguna-api/database/seeders"
	"depoguna-api/handlers"
	"depoguna-api/models"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load env file")
	}
}

func main() {
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// migration
	migrations.AutoMigration(db)

	// seeder to create customer
	if err := db.AutoMigrate(&models.Customer{}); err == nil && db.Migrator().HasTable(&models.Customer{}) {
		if err := db.First(&models.Customer{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeders.CreateCustomer(db)
		}
	}

	r := gin.Default()

	// customer handler
	customerHandler := handlers.NewCustomerHandler(db)
	customerHandler.Handler(&r.RouterGroup)
	r.Run(":8080")
}
