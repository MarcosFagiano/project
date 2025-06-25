// package tests
//
// import (
//
//	"backend/models"
//	"backend/utils"
//	"gorm.io/gorm"
//	"log"
//	"time"
//
// )
//
//	func hashPassword(password string) string {
//		return utils.HashPassword(password)
//	}
//
// func SeedTestData(db *gorm.DB) error {
//
//		// Crear categorías
//
//		// Hashea la contraseña en SHA256
//
//		categories := []models.Category{
//			{Name: "Funcional", Description: "Entrenamiento funcional"},
//			{Name: "Spinning", Description: "Clases de bicicleta fija"},
//			{Name: "MMA", Description: "Artes marciales mixtas"},
//			{Name: "Yoga", Description: "Clases de yoga"},
//			{Name: "Pilates", Description: "Pilates"},
//			{Name: "Crossfit", Description: "Crossfit"},
//			{Name: "Natacion", Description: "Natación libre"},
//			{Name: "Zumba", Description: "Clases de zumba"},
//			{Name: "Boxeo", Description: "Boxeo para principiantes"},
//			{Name: "HIIT", Description: "Entrenamiento HIIT"},
//			{Name: "Tai Chi", Description: "Tai Chi básico"},
//			{Name: "Escalada", Description: "Indoor climbing"},
//			{Name: "Karate", Description: "Karate do"},
//			{Name: "Ciclismo", Description: "Ciclismo indoor"},
//			{Name: "Remo", Description: "Remo fijo"},
//			{Name: "Básquet", Description: "Entrenamiento de básquet"},
//			{Name: "Fútbol", Description: "Fútbol 5"},
//			{Name: "Handball", Description: "Handball recreativo"},
//			{Name: "Tenis", Description: "Tenis nivel inicial"},
//			{Name: "Voley", Description: "Voley mixto"},
//		}
//		if err := db.Create(&categories).Error; err != nil {
//			return err
//		}
//
//		// Crear usuarios
//		users := []models.User{
//			{GivenName: "Juan", FamilyName: "Perez", PhoneNumber: 1234, Email: "juan@example.com", Admin: false, Password: hashPassword("pass1")},
//			{GivenName: "Ana", FamilyName: "Lopez", PhoneNumber: 5678, Email: "ana@example.com", Admin: false, Password: hashPassword("pass2")},
//			{GivenName: "Luis", FamilyName: "Martinez", PhoneNumber: 4321, Email: "luis@example.com", Admin: false, Password: hashPassword("pass3")},
//			{GivenName: "Carlos", FamilyName: "Gomez", PhoneNumber: 1111, Email: "carlos@example.com", Admin: false, Password: hashPassword("pass4")},
//			{GivenName: "Laura", FamilyName: "Diaz", PhoneNumber: 2222, Email: "laura@example.com", Admin: false, Password: hashPassword("pass5")},
//			{GivenName: "Pedro", FamilyName: "Suarez", PhoneNumber: 3333, Email: "pedro@example.com", Admin: false, Password: hashPassword("pass6")},
//			{GivenName: "Marta", FamilyName: "Fernandez", PhoneNumber: 4444, Email: "marta@example.com", Admin: false, Password: hashPassword("pass7")},
//			{GivenName: "Jose", FamilyName: "Garcia", PhoneNumber: 5555, Email: "jose@example.com", Admin: false, Password: hashPassword("pass8")},
//			{GivenName: "Sofia", FamilyName: "Ramirez", PhoneNumber: 6666, Email: "sofia@example.com", Admin: false, Password: hashPassword("pass9")},
//			{GivenName: "Lucia", FamilyName: "Mendez", PhoneNumber: 7777, Email: "lucia@example.com", Admin: false, Password: hashPassword("pass10")},
//			{GivenName: "Diego", FamilyName: "Alvarez", PhoneNumber: 8888, Email: "diego@example.com", Admin: false, Password: hashPassword("pass11")},
//			{GivenName: "Paula", FamilyName: "Torres", PhoneNumber: 9999, Email: "paula@example.com", Admin: false, Password: hashPassword("pass12")},
//			{GivenName: "Andres", FamilyName: "Vega", PhoneNumber: 1010, Email: "andres@example.com", Admin: false, Password: hashPassword("pass13")},
//			{GivenName: "Nico", FamilyName: "Ortega", PhoneNumber: 2020, Email: "nico@example.com", Admin: false, Password: hashPassword("pass14")},
//			{GivenName: "Valeria", FamilyName: "Sosa", PhoneNumber: 3030, Email: "valeria@example.com", Admin: false, Password: hashPassword("pass15")},
//			{GivenName: "Ricardo", FamilyName: "Luna", PhoneNumber: 4040, Email: "ricardo@example.com", Admin: false, Password: hashPassword("pass16")},
//			{GivenName: "Cecilia", FamilyName: "Paz", PhoneNumber: 5050, Email: "cecilia@example.com", Admin: false, Password: hashPassword("pass17")},
//			{GivenName: "Mario", FamilyName: "Vargas", PhoneNumber: 6060, Email: "mario@example.com", Admin: false, Password: hashPassword("pass18")},
//			{GivenName: "Flor", FamilyName: "Reyes", PhoneNumber: 7070, Email: "flor@example.com", Admin: false, Password: hashPassword("pass19")},
//			{GivenName: "God", FamilyName: "Admin", PhoneNumber: 666, Email: "god@admin.com", Admin: true, Password: hashPassword("holyfather")},
//		}
//		if err := db.Create(&users).Error; err != nil {
//			return err
//		}
//
//		// Crear actividades
//		activities := []models.Activity{
//			{Description: "Funcional Lunes", StartDate: parseDateTime("2025-07-01 10:00:00"), EndDate: parseDateTime("2025-07-01 11:00:00"), Quota: 20, CategoryID: 1},
//			{Description: "Spinning Martes", StartDate: parseDateTime("2025-07-02 09:00:00"), EndDate: parseDateTime("2025-07-02 10:00:00"), Quota: 15, CategoryID: 2},
//			{Description: "MMA Miércoles", StartDate: parseDateTime("2025-07-03 18:00:00"), EndDate: parseDateTime("2025-07-03 19:30:00"), Quota: 10, CategoryID: 3},
//			{Description: "Yoga Jueves", StartDate: parseDateTime("2025-07-04 08:00:00"), EndDate: parseDateTime("2025-07-04 09:00:00"), Quota: 25, CategoryID: 4},
//			{Description: "Pilates Viernes", StartDate: parseDateTime("2025-07-05 17:00:00"), EndDate: parseDateTime("2025-07-05 18:00:00"), Quota: 20, CategoryID: 5},
//			{Description: "Crossfit Sábado", StartDate: parseDateTime("2025-07-06 19:00:00"), EndDate: parseDateTime("2025-07-06 20:00:00"), Quota: 15, CategoryID: 6},
//			{Description: "Zumba Domingo", StartDate: parseDateTime("2025-07-07 11:00:00"), EndDate: parseDateTime("2025-07-07 12:00:00"), Quota: 30, CategoryID: 8},
//			{Description: "Boxeo Lunes", StartDate: parseDateTime("2025-07-08 18:00:00"), EndDate: parseDateTime("2025-07-08 19:00:00"), Quota: 12, CategoryID: 9},
//			{Description: "HIIT Martes", StartDate: parseDateTime("2025-07-09 07:00:00"), EndDate: parseDateTime("2025-07-09 07:45:00"), Quota: 18, CategoryID: 10},
//			{Description: "Tai Chi Miércoles", StartDate: parseDateTime("2025-07-10 06:30:00"), EndDate: parseDateTime("2025-07-10 07:30:00"), Quota: 20, CategoryID: 11},
//		}
//		if err := db.Create(&activities).Error; err != nil {
//			return err
//		}
//
//		// Crear inscripciones
//		inscriptions := []models.Inscription{
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 1},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 2},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 3},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 4},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 5},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 6},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 7},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 8},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 9},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 10},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 11},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 12},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 13},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 14},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 15},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 16},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 17},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 18},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 19},
//			{StartDate: "2025-07-01", EndDate: "2025-07-31", UserID: 20},
//		}
//		if err := db.Create(&inscriptions).Error; err != nil {
//			return err
//		}
//
//		// Relacionar inscripciones y actividades
//		for i, insc := range inscriptions {
//			a1 := activities[i%len(activities)]
//			a2 := activities[(i+1)%len(activities)]
//
//			if err := db.Model(&insc).Association("Activities").Append(&a1, &a2); err != nil {
//				log.Printf("Error asociando actividades: %v", err)
//			}
//		}
//
//		return nil
//	}
//
//	func parseDateTime(s string) time.Time {
//		t, _ := time.Parse("2006-01-02 15:04:05", s)
//		return t
//	}
package tests

import (
	"backend/models"
	"backend/utils"
	"gorm.io/gorm"
	"time"
)

func hashPassword(password string) string {
	return utils.HashPassword(password)
}

func SeedTestData(db *gorm.DB) error {
	//// Crear categorías
	//categories := []models.Category{
	//	{Name: "Funcional", Description: "Entrenamiento funcional"},
	//	{Name: "Spinning", Description: "Clases de bicicleta fija"},
	//	{Name: "MMA", Description: "Artes marciales mixtas"},
	//	{Name: "Yoga", Description: "Clases de yoga"},
	//	{Name: "Pilates", Description: "Pilates"},
	//	{Name: "Crossfit", Description: "Crossfit"},
	//	{Name: "Natacion", Description: "Natación libre"},
	//	{Name: "Zumba", Description: "Clases de zumba"},
	//	{Name: "Boxeo", Description: "Boxeo para principiantes"},
	//	{Name: "HIIT", Description: "Entrenamiento HIIT"},
	//	{Name: "Tai Chi", Description: "Tai Chi básico"},
	//	{Name: "Escalada", Description: "Indoor climbing"},
	//	{Name: "Karate", Description: "Karate do"},
	//	{Name: "Ciclismo", Description: "Ciclismo indoor"},
	//	{Name: "Remo", Description: "Remo fijo"},
	//	{Name: "Básquet", Description: "Entrenamiento de básquet"},
	//	{Name: "Fútbol", Description: "Fútbol 5"},
	//	{Name: "Handball", Description: "Handball recreativo"},
	//	{Name: "Tenis", Description: "Tenis nivel inicial"},
	//	{Name: "Voley", Description: "Voley mixto"},
	//}
	//if err := db.Create(&categories).Error; err != nil {
	//	return err
	//}

	//// Crear usuarios
	//users := []models.User{
	//	{GivenName: "Juan", FamilyName: "Perez", PhoneNumber: "1234", Email: "juan@example.com", Admin: false, Password: hashPassword("pass1")},
	//	{GivenName: "Ana", FamilyName: "Lopez", PhoneNumber: "5678", Email: "ana@example.com", Admin: false, Password: hashPassword("pass2")},
	//	{GivenName: "Luis", FamilyName: "Martinez", PhoneNumber: "4321", Email: "luis@example.com", Admin: false, Password: hashPassword("pass3")},
	//	{GivenName: "Carlos", FamilyName: "Gomez", PhoneNumber: "1111", Email: "carlos@example.com", Admin: false, Password: hashPassword("pass4")},
	//	{GivenName: "Laura", FamilyName: "Diaz", PhoneNumber: "2222", Email: "laura@example.com", Admin: false, Password: hashPassword("pass5")},
	//	{GivenName: "Pedro", FamilyName: "Suarez", PhoneNumber: "3333", Email: "pedro@example.com", Admin: false, Password: hashPassword("pass6")},
	//	{GivenName: "Marta", FamilyName: "Fernandez", PhoneNumber: "4444", Email: "marta@example.com", Admin: false, Password: hashPassword("pass7")},
	//	{GivenName: "Jose", FamilyName: "Garcia", PhoneNumber: "5555", Email: "jose@example.com", Admin: false, Password: hashPassword("pass8")},
	//	{GivenName: "Sofia", FamilyName: "Ramirez", PhoneNumber: "6666", Email: "sofia@example.com", Admin: false, Password: hashPassword("pass9")},
	//	{GivenName: "Lucia", FamilyName: "Mendez", PhoneNumber: "7777", Email: "lucia@example.com", Admin: false, Password: hashPassword("pass10")},
	//	{GivenName: "Diego", FamilyName: "Alvarez", PhoneNumber: "8888", Email: "diego@example.com", Admin: false, Password: hashPassword("pass11")},
	//	{GivenName: "Paula", FamilyName: "Torres", PhoneNumber: "9999", Email: "paula@example.com", Admin: false, Password: hashPassword("pass12")},
	//	{GivenName: "Andres", FamilyName: "Vega", PhoneNumber: "1010", Email: "andres@example.com", Admin: false, Password: hashPassword("pass13")},
	//	{GivenName: "Nico", FamilyName: "Ortega", PhoneNumber: "2020", Email: "nico@example.com", Admin: false, Password: hashPassword("pass14")},
	//	{GivenName: "Valeria", FamilyName: "Sosa", PhoneNumber: "3030", Email: "valeria@example.com", Admin: false, Password: hashPassword("pass15")},
	//	{GivenName: "Ricardo", FamilyName: "Luna", PhoneNumber: "4040", Email: "ricardo@example.com", Admin: false, Password: hashPassword("pass16")},
	//	{GivenName: "Cecilia", FamilyName: "Paz", PhoneNumber: "5050", Email: "cecilia@example.com", Admin: false, Password: hashPassword("pass17")},
	//	{GivenName: "Mario", FamilyName: "Vargas", PhoneNumber: "6060", Email: "mario@example.com", Admin: false, Password: hashPassword("pass18")},
	//	{GivenName: "Flor", FamilyName: "Reyes", PhoneNumber: "7070", Email: "flor@example.com", Admin: false, Password: hashPassword("pass19")},
	//	{GivenName: "God", FamilyName: "Admin", PhoneNumber: "666", Email: "god@admin.com", Admin: true, Password: hashPassword("holyfather")},
	//}
	//
	//if err := db.Create(&users).Error; err != nil {
	//	return err
	//}

	activities := []models.Activity{
		{Description: "Funcional Lunes", StartDate: parseDateTime("2025-07-01 10:00:00"), EndDate: parseDateTime("2025-07-01 11:00:00"), Quota: 20, CategoryID: 1},
		{Description: "Spinning Martes", StartDate: parseDateTime("2025-07-02 09:00:00"), EndDate: parseDateTime("2025-07-02 10:00:00"), Quota: 15, CategoryID: 2},
		{Description: "MMA Miércoles", StartDate: parseDateTime("2025-07-03 18:00:00"), EndDate: parseDateTime("2025-07-03 19:30:00"), Quota: 10, CategoryID: 3},
		{Description: "Yoga Jueves", StartDate: parseDateTime("2025-07-04 08:00:00"), EndDate: parseDateTime("2025-07-04 09:00:00"), Quota: 25, CategoryID: 4},
		{Description: "Pilates Viernes", StartDate: parseDateTime("2025-07-05 17:00:00"), EndDate: parseDateTime("2025-07-05 18:00:00"), Quota: 20, CategoryID: 5},
		{Description: "Crossfit Sábado", StartDate: parseDateTime("2025-07-06 19:00:00"), EndDate: parseDateTime("2025-07-06 20:00:00"), Quota: 15, CategoryID: 6},
		{Description: "Zumba Domingo", StartDate: parseDateTime("2025-07-07 11:00:00"), EndDate: parseDateTime("2025-07-07 12:00:00"), Quota: 30, CategoryID: 8},
		{Description: "Boxeo Lunes", StartDate: parseDateTime("2025-07-08 18:00:00"), EndDate: parseDateTime("2025-07-08 19:00:00"), Quota: 12, CategoryID: 9},
		{Description: "HIIT Martes", StartDate: parseDateTime("2025-07-09 07:00:00"), EndDate: parseDateTime("2025-07-09 07:45:00"), Quota: 18, CategoryID: 10},
		{Description: "Tai Chi Miércoles", StartDate: parseDateTime("2025-07-10 06:30:00"), EndDate: parseDateTime("2025-07-10 07:30:00"), Quota: 20, CategoryID: 11},
	}
	if err := db.Create(&activities).Error; err != nil {
		return err
	}

	return nil
}

func parseDateTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return t
}
