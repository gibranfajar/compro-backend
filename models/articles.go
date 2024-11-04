package models

import (
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	Title       string     `gorm:"not null" json:"title" form:"title"`
	Description string     `gorm:"not null" json:"description" form:"description"`
	CategoryID  uint       `gorm:"not null" json:"category_id" form:"category_id"`
	Category    Categories `gorm:"foreignKey:CategoryID"`
}
