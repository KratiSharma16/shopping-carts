package routes

import (
    "shopping-cart/controllers"
    "shopping-cart/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/users", controllers.CreateUser)
    r.POST("/users/login", controllers.LoginUser)
    r.GET("/users", controllers.GetUsers)

    r.POST("/items", controllers.CreateItem)
    r.GET("/items", controllers.GetItems)

    auth := r.Group("/")
    auth.Use(middleware.JWTAuth())

    auth.POST("/carts", controllers.AddToCart)
    auth.GET("/carts", controllers.GetCarts)

    auth.POST("/orders", controllers.CreateOrder)
    auth.GET("/orders", controllers.GetOrders)
}
