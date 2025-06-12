package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("backend.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	DB = database
	log.Println("Conexión a la base de datos establecida.")
}
