package models

import "time"

type Inscription struct {
	ID         uint
	UserID     uint
	ActivityID uint `gorm:"foreignKey:ActivityID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
