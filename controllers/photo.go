package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/app"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/config"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/helpers"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/models"
	"gorm.io/gorm"
	"net/http"
)

var (
	errPhotoNotFound = errors.New("photo not found")
)

type photoController struct {
	serverConfig config.Config
	database     *gorm.DB
}

// NewPhotoController create instance of photo controller
func NewPhotoController(database *gorm.DB, serverConfig config.Config) *photoController {
	pc := photoController{
		serverConfig: serverConfig,
		database:     database,
	}

	return &pc
}

func (pc *photoController) POSTInsertPhoto(c *gin.Context) {
	tokenClaims, ok := c.MustGet("claims").(helpers.Claims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	var req app.InsertPhotoRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.CustomValidationErrorMessage(err.(validator.ValidationErrors)))

		return
	}

	photo := models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   tokenClaims.UserID,
	}

	if err := pc.database.Model(models.Photo{}).Create(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, app.InsertPhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	})
}

func (pc *photoController) GETFindAllPhotos(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (pc *photoController) PUTUpdatePhoto(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (pc *photoController) DELETEDeletePhoto(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
