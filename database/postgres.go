package database

import (
	"fmt"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func Start(config Config, runMigration bool, logger *logrus.Logger) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)))
	if err != nil {
		panic(err)
	}

	logger.Infoln("[~] database connected")

	if runMigration {
		err = conn.AutoMigrate(&models.User{}, &models.Photo{})
		if err != nil {
			panic(err)
		}

		logger.Infoln("[~] database migration successfully")
	}

	return conn
}
