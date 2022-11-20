package seeders

import (
	"depoguna-api/models"

	"gorm.io/gorm"
)

// CreateOrder is seeder to create order
func CreateOrder(db *gorm.DB) error {
	var orders = []models.Order{
		{CustomerId: 1, ProductId: 2, Qty: 20},
		{CustomerId: 2, ProductId: 3, Qty: 10},
		{CustomerId: 3, ProductId: 3, Qty: 30},
		{CustomerId: 4, ProductId: 2, Qty: 40},
		{CustomerId: 5, ProductId: 2, Qty: 50},
		{CustomerId: 6, ProductId: 2, Qty: 15},
		{CustomerId: 7, ProductId: 1, Qty: 45},
		{CustomerId: 8, ProductId: 4, Qty: 25},
		{CustomerId: 9, ProductId: 5, Qty: 35},
		{CustomerId: 10, ProductId: 7, Qty: 65},
		{CustomerId: 11, ProductId: 8, Qty: 45},
		{CustomerId: 12, ProductId: 5, Qty: 35},
		{CustomerId: 13, ProductId: 3, Qty: 25},
		{CustomerId: 14, ProductId: 2, Qty: 45},
		{CustomerId: 15, ProductId: 1, Qty: 15},
	}
	return db.Create(&orders).Error
}
