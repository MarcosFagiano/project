package models

import "time"

type User struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	GivenName         string    `gorm:"size:100;not null"`
	FamilyName        string    `gorm:"size:100;not null"`
	PhoneNumber       string    `gorm:"size:20"`
	Email             string    `gorm:"size:100;unique;not null"`
	Admin             bool      `gorm:"default:false;not null"`
	Password          string    `gorm:"not null"`
	RegisterDateStart time.Time `gorm:"not null"`
	RegisterDateEnd   time.Time `gorm:"not null"`
}
