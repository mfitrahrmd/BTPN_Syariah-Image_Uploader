package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/config"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/controllers"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/database"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/middlewares"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/router"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	ServerConfig   *config.Config
	RouterEngine   *gin.Engine
	DatabaseEngine *gorm.DB
	Logger         *logrus.Logger
}

func BuildServer() *Server {
	routerInstance := gin.Default()
	loggerInstance := logrus.New()
	configInstance := config.New(".", loggerInstance)
	databaseInstance := database.Start(database.Config{
		User:     configInstance.PostgresUser,
		Password: configInstance.PostgresPassword,
		Host:     configInstance.PostgresHost,
		Port:     configInstance.PostgresPort,
		Database: configInstance.PostgresDb,
	}, true, loggerInstance)

	serverInstance := Server{
		ServerConfig:   configInstance,
		RouterEngine:   routerInstance,
		DatabaseEngine: databaseInstance,
		Logger:         loggerInstance,
	}

	serverInstance.RouterEngine.Use(middlewares.ErrorHandler(serverInstance.Logger))
	router.WithUserRoutes(serverInstance.RouterEngine, controllers.NewUserController(serverInstance.DatabaseEngine, serverInstance.ServerConfig))
	router.WithPhotoRoutes(serverInstance.RouterEngine, controllers.NewPhotoController(serverInstance.DatabaseEngine, serverInstance.ServerConfig), middlewares.NewAuthMiddleware(serverInstance.DatabaseEngine, serverInstance.ServerConfig), middlewares.NewPhotoMiddleware(serverInstance.DatabaseEngine, serverInstance.ServerConfig))

	return &serverInstance
}

func (s *Server) Run() error {
	return s.RouterEngine.Run(fmt.Sprintf("%s:%s", s.ServerConfig.AppHost, s.ServerConfig.AppPort))
}
