package models

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	GivenName   string `gorm:"size:100;not null"`
	FamilyName  string `gorm:"size:100;not null"`
	PhoneNumber string `gorm:"size:20"`
	Email       string `gorm:"size:100;unique;not null"`
	Password    string `gorm:"not null"`
	Admin       bool   `gorm:"default:false"`
	//InscriptionID uint          // Clave foránea
	Inscription []Inscription `gorm:"foreignKey:UserID"`
}
