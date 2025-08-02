package models

import "github.com/jinzhu/gorm"

type Item struct {
    gorm.Model
    Name  string
    Price float64
}
