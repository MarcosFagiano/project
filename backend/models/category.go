package models

type Category struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:100;unique;not null"`
	Description string `gorm:"size:255"`
}
