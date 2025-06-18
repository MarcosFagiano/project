package models

type Inscription struct {
	ID         uint        `gorm:"primaryKey;autoIncrement"`
	StartDate  string      `gorm:"type:date"`
	EndDate    string      `gorm:"type:date"`
	Activities []*Activity `gorm:"many2many:inscription_activities"`
	UserID     uint
}
