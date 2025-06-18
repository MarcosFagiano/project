package main

import (
	"backend/config"
	"backend/routes"
	"log"
)

func main() {
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = routes.Init()
	if err != nil {
		log.Fatal(err)
	}

}
