package controllers

import (
    "shopping-cart/config"
    "shopping-cart/models"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "time"
)

var jwtKey = []byte("secret_key")

func CreateUser(c *gin.Context) {
    var user models.User
    c.BindJSON(&user)
    config.DB.Create(&user)
    c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
    var input models.User
    c.BindJSON(&input)

    var user models.User
    config.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user)

    if user.ID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenString, _ := token.SignedString(jwtKey)
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
func GetUsers(c *gin.Context) {
    var users []models.User
    config.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}