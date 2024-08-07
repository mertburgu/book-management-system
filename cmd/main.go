package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mertburgu/book-management-system/api"
    "github.com/mertburgu/book-management-system/config"
)

func main() {
    config.LoadConfig()
    config.InitDatabase() // Initialize the database

    r := gin.Default()

    // Register routes
    api.RegisterRoutes(r)

    r.Run(":8082") // Start server
}
