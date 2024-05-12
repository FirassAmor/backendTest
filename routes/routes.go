package routes

import (
    "github.com/gin-gonic/gin"
    "backendtest/controllers"
)

func InitializeRoutes(router *gin.Engine) {
    router.GET("/person/:person_id/info", controllers.GetPersonInfo)
    router.POST("/person/create", controllers.CreatePerson)

    router.GET("/address/:address_id/info", controllers.GetAddressInfo)
}
