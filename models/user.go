package models

import "time"

type User struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp(0);autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);autoUpdateTime"`
}
