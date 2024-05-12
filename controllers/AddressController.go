package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "backendtest/services"
)

func GetAddressInfo(c *gin.Context) {
    addressID := c.Param("address_id")
    info, err := services.GetAddressInfo(addressID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, info)
}
