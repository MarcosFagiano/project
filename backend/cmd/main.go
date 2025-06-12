package main

import (
    "myapp/config"
    "myapp/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
