package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `gorm:"type:string"`
	Caption  string `gorm:"type:string"`
	PhotoUrl string `gorm:"type:string;not null;unique"`
	UserID   uint
}
