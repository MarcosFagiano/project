package main

import (
	"backend/config"
	"backend/routes"
	"log"
)

func main() {
	//config.LoadEnv()
	//r := gin.Default()
	//routes.SetupRoutes(r)
	//r.Run(":8080")
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = routes.Init()
	if err != nil {
		log.Fatal(err)
	}

}
