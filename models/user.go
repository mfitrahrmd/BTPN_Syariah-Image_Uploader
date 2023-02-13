package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"type:string;size:16;not null;unique"`
	Email    string  `gorm:"type:string;size:50;not null;unique;index"`
	Password string  `gorm:"type:string;not null"`
	Photos   []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
