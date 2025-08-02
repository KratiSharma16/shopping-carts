package controllers

import (
    "net/http"
    "shopping-cart/config"
    "shopping-cart/models"

    "github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    userID := getUserIDFromToken(c)

    var cart models.Cart
    result := config.DB.Where("user_id = ?", userID).First(&cart)
    if result.Error != nil || cart.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No cart found"})
        return
    }

    order := models.Order{UserID: userID, CartID: cart.ID}
    config.DB.Create(&order)

    c.JSON(http.StatusOK, gin.H{"message": "Order placed", "order_id": order.ID})
}

func GetOrders(c *gin.Context) {
    userID := getUserIDFromToken(c)

    var orders []models.Order
    config.DB.Where("user_id = ?", userID).Find(&orders)

    c.JSON(http.StatusOK, orders)
}
