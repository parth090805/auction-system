package main

import (
    "auction-system/config"
    "auction-system/models"
    "auction-system/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.Connect()
    config.DB.AutoMigrate(&models.User{})

    r := gin.Default()
    routes.RegisterRoutes(r)
    r.Run(":8080")
}
