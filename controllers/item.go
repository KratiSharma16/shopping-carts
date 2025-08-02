package controllers

import (
    "net/http"
    "shopping-cart/config"
    "shopping-cart/models"

    "github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
    var item models.Item
    if err := c.BindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&item)
    c.JSON(http.StatusOK, item)
}

func GetItems(c *gin.Context) {
    var items []models.Item
    config.DB.Find(&items)
    c.JSON(http.StatusOK, items)
}
