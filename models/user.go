package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `json:"username,omitempty" gorm:"type:string;size:16;not null;unique"`
	Email    string  `json:"email,omitempty" gorm:"type:string;size:50;not null;unique;index"`
	Password string  `json:"password,omitempty" gorm:"type:string;not null"`
	Photos   []Photo `json:"photos,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
