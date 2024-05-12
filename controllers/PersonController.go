package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
	"strconv"
    "backendtest/services"
	"backendtest/models"
)

func GetPersonInfo(c *gin.Context) {
    personID := c.Param("person_id")
    info, err := services.GetPersonInfo(personID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, info)
}

func CreatePerson(c *gin.Context) {
    var person models.Person

    if err := c.BindJSON(&person); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.CreatePerson(&person)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Person created successfully"})
}

func GetAllPersonsInfo(c *gin.Context) {
    // Retrieve pagination parameters
    limit := c.DefaultQuery("limit", "10") // Default limit of 10 if not provided
    offset := c.DefaultQuery("offset", "0")

    // Convert limit and offset to integer
    limitInt, _ := strconv.Atoi(limit)
    offsetInt, _ := strconv.Atoi(offset)

    // Call the service function with pagination parameters
    info, err := services.GetAllPersonInfo(limitInt, offsetInt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, info)
}