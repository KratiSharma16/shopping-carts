package models

import "github.com/jinzhu/gorm"

type Cart struct {
    gorm.Model
    UserID uint
    Items  []CartItem
}

type CartItem struct {
    gorm.Model
    CartID uint
    ItemID uint
}
