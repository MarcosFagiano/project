package models

import "time"

type Activity struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"size:255"`
	StartDate   time.Time
	EndDate     time.Time
	Quota       int `gorm:"type:int DEFAULT 15"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
}
