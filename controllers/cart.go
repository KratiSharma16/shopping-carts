package controllers

import (
    "net/http"
    "shopping-cart/config"
    "shopping-cart/models"

    "github.com/gin-gonic/gin"
)

type AddToCartInput struct {
    ItemID uint `json:"item_id"`
}

func AddToCart(c *gin.Context) {
    var input AddToCartInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Get user ID from token claims (optional improvement)
    userID := getUserIDFromToken(c) // you can extract this from JWT if implemented

    // Find or create cart for user
    var cart models.Cart
    config.DB.Where("user_id = ?", userID).FirstOrCreate(&cart)

    cartItem := models.CartItem{CartID: cart.ID, ItemID: input.ItemID}
    config.DB.Create(&cartItem)

    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func GetCarts(c *gin.Context) {
    userID := getUserIDFromToken(c)

    var carts []models.Cart
    config.DB.Preload("Items").Where("user_id = ?", userID).Find(&carts)

    c.JSON(http.StatusOK, carts)
}

// Dummy function – replace with real JWT parsing
func getUserIDFromToken(c *gin.Context) uint {
    // In real project, extract from JWT claims in middleware
    return 1 // assuming user with ID 1 is logged in for now
}
