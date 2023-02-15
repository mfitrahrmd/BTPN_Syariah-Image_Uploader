package database

import (
	"fmt"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func StartDatabase(config Config, runMigration bool) (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)))
	if err != nil {
		return nil, fmt.Errorf("err connecting to database : %w", err)
	}

	logrus.Infoln("[~] database connected")

	if runMigration {
		err = conn.AutoMigrate(&models.User{}, &models.Photo{})
		if err != nil {
			return nil, fmt.Errorf("err running database migration : %w", err)
		}

		log.Println("[~] database migration successfully")
	}

	return conn, nil
}
