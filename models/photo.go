package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `json:"title,omitempty" gorm:"type:string;not null"`
	Caption  string `json:"caption,omitempty" gorm:"type:string;not null"`
	PhotoUrl string `json:"photoUrl,omitempty" gorm:"type:string;not null"`
	UserID   uint   `json:"userID,omitempty" gorm:"not null"`
}
