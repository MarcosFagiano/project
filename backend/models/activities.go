package models

import "time"

type Activity struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	CategoryID  uint `gorm:"foreignKey:CategoryID"`
	StartDate   time.Time
	EndDate     time.Time
	Quota       uint
	Description string `gorm:"size:255"`
}
