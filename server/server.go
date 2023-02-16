package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/config"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/controllers"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/database"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/middlewares"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/router"
	"gorm.io/gorm"
)

type Server struct {
	ServerConfig config.Config
	RouterEngine *gin.Engine
	Repository   *gorm.DB
}

func BuildServer() (*Server, error) {
	var s Server

	s.RouterEngine = gin.Default()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		return nil, fmt.Errorf("err loading config : %w", err)
	}

	s.ServerConfig = cfg

	db, err := database.StartDatabase(database.Config{
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		Database: cfg.PostgresDb,
	}, true)
	if err != nil {
		return nil, fmt.Errorf("err starting database : %w", err)
	}

	s.Repository = db

	router.WithUserRoutes(s.RouterEngine, controllers.NewUserController(s.Repository, s.ServerConfig))
	router.WithPhotoRoutes(s.RouterEngine, controllers.NewPhotoController(s.Repository, s.ServerConfig), middlewares.NewAuthMiddleware(s.Repository, s.ServerConfig), middlewares.NewPhotoMiddleware(s.Repository, s.ServerConfig))

	return &s, nil
}

func (s *Server) Run() error {
	return s.RouterEngine.Run(fmt.Sprintf("%s:%s", s.ServerConfig.AppHost, s.ServerConfig.AppPort))
}
