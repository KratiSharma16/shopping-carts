package models

import "github.com/jinzhu/gorm"

type Order struct {
    gorm.Model
    UserID uint
    CartID uint
}
