package config

import (
	"backend/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	log.Printf("Establish conection to DB in %s:%s/%s", dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}
	DB = database

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Activity{},
		&models.Category{},
		&models.Inscription{},
		// &models.InscriptionActivity{}, // si existe explícitamente
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
		return nil, err
	}
	return DB, nil
}
