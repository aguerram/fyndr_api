package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint64 `gorm:"primarykey;autoIncrement:true"`
	Name string `gorm:"unique;not null"`
}
