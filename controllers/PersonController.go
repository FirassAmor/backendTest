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
    offset, _ := strconv.Atoi(c.Query("offset"))
    limit, _ := strconv.Atoi(c.Query("limit"))

    persons, totalCount, err := services.GetAllPersons(offset, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"persons": persons, "total": totalCount})
}