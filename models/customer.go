package models

import "time"

type Customer struct {
	Id          int       `gorm:"primaryKey;autoIncrement"`
	UserId      int       `gorm:"not null"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"type:varchar(255);not null"`
	Gender      string    `gorm:"type:varchar(255);not null"`
	DateOfBirth string    `gorm:"type:varchar(255);not null"`
	Mobile      string    `gorm:"type:varchar(255);not null"`
	Address     string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"type:timestamp(0);autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp(0);autoUpdateTime"`
}
