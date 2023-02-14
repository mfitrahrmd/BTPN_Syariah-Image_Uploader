package database

import (
	"fmt"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func RunDatabase(config Config, runMigration bool) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)))
	if err != nil {
		panic(fmt.Sprintf("err connecting to database : %v\n", err))
	}

	if runMigration {
		err = conn.AutoMigrate(&models.User{}, &models.Photo{})
		if err != nil {
			panic(fmt.Sprintf("err running database migration : %v\n", err))
		}
	}

	return conn
}
