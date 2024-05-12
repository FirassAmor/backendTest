package main

import (
    "github.com/gin-gonic/gin"
    "backendtest/routes"
    "backendtest/database"
)

func main() {
    db, err := database.Connect()
    if err != nil {
        panic(err)
    }
    defer db.Close()

    r := gin.Default()

    routes.InitializeRoutes(r)

    r.Run(":8080")
}
