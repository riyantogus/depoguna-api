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

	// seeder to create user
	if err := db.AutoMigrate(&models.User{}); err == nil && db.Migrator().HasTable(&models.User{}) {
		if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeders.CreateUser(db)
		}
	}

	// seeder to create customer
	if err := db.AutoMigrate(&models.Customer{}); err == nil && db.Migrator().HasTable(&models.Customer{}) {
		if err := db.First(&models.Customer{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeders.CreateCustomer(db)
		}
	}

	// seeder to create order
	if err := db.AutoMigrate(&models.Order{}); err == nil && db.Migrator().HasTable(&models.Order{}) {
		if err := db.First(&models.Order{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeders.CreateOrder(db)
		}
	}

	r := gin.Default()

	// auth handler
	authHandler := handlers.NewAuthHandler(db)
	authHandler.Handler(&r.RouterGroup)

	// customer handler
	customerHandler := handlers.NewCustomerHandler(db)
	customerHandler.Handler(&r.RouterGroup)

	// order handler
	orderHandler := handlers.NewOrderHandler(db)
	orderHandler.Handler(&r.RouterGroup)

	r.Run(":8080")
}
