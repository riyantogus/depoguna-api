package models

import "time"

type Order struct {
	Id         int       `gorm:"primaryKey;autoIncrement"`
	CustomerId int       `gorm:"not null"`
	ProductId  int       `gorm:"not null"`
	Qty        int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"type:timestamp(0);autoCreateTime"`
	UpdatedAt  time.Time `gorm:"type:timestamp(0);autoUpdateTime"`
}
