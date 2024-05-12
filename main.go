package main

import (
    "github.com/gin-gonic/gin"
    "backendtest/routes"
    "backendtest/database"
    "net/http"
)

func main() {
    db, err := database.Connect()
    if err != nil {
        panic(err)
    }
    defer db.Close()

    r := gin.Default()

    // Use the CORS middleware directly in Gin's Use function
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }

        c.Next()
    })

    // Initialize routes
    routes.InitializeRoutes(r)

    // Run the Gin application
    r.Run(":8080")
}
