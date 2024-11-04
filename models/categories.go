package models

import (
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string `gorm:"not null unique" json:"name" form:"name"`
}
