package main

import (
	"backend/config"
	"backend/dependencies"
	"backend/routes"
	"log"
)

func main() {

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	// Inyeccion de dependecias
	dep, err := dependencies.BuildDependencies(db)
	if err != nil {
		log.Fatal(err)
	}
	err = routes.Init(dep)
	if err != nil {
		log.Fatal(err)
	}
	//err = tests.SeedTestData(config.DB)
	//if err != nil {
	//	return
	//} //load of test data
}
