package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/app"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/config"
	controllerError "github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/error"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/helpers"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/models"
	"gorm.io/gorm"
	"net/http"
)

var (
	errPhotoNotFound    = errors.New("photo not found")
	errTokenClaimsEmpty = errors.New("token claims empty")
)

type photoController struct {
	serverConfig *config.Config
	database     *gorm.DB
}

// NewPhotoController create instance of photo controller
func NewPhotoController(database *gorm.DB, serverConfig *config.Config) *photoController {
	pc := photoController{
		serverConfig: serverConfig,
		database:     database,
	}

	return &pc
}

func (pc *photoController) POSTInsertPhoto(c *gin.Context) {
	tokenClaims, ok := c.MustGet("claims").(helpers.Claims)
	if !ok {
		c.Error(errTokenClaimsEmpty)

		return
	}

	var req app.InsertPhotoRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(controllerError.New(http.StatusBadRequest, errValidation.Error(), helpers.CustomValidationErrorMessage(err.(validator.ValidationErrors))))

		return
	}

	photo := models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   tokenClaims.UserID,
	}

	if err := pc.database.Model(models.Photo{}).Create(&photo).Error; err != nil {
		c.Error(err)

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
	var res app.FindAllPhotosResponse

	if err := pc.database.Model(&models.Photo{}).Find(&res.Photos).Error; err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, res)
}

func (pc *photoController) PUTUpdatePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	var req app.UpdatePhotoRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(controllerError.New(http.StatusBadRequest, errValidation.Error(), helpers.CustomValidationErrorMessage(err.(validator.ValidationErrors))))

		return
	}

	// check if photo to be updated is exists in repository
	var photo models.Photo

	if err = pc.database.Model(&models.Photo{}).First(&photo, photoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(controllerError.New(http.StatusNotFound, errPhotoNotFound.Error()))

			return
		}
		c.Error(err)

		return
	}

	// update the photo data
	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoUrl = req.PhotoUrl

	if err = pc.database.Model(&models.Photo{}).Where("id = ?", photoId).Updates(&photo).Error; err != nil {
		c.Error(err)

		return
	}

	// send response with updated photo's data
	c.JSON(http.StatusOK, app.UpdatePhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	})
}

func (pc *photoController) DELETEDeletePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	// check if photo to be updated is exists in repository
	var photo models.Photo

	if err := pc.database.Model(&models.Photo{}).First(&photo, photoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(controllerError.New(http.StatusNotFound, errPhotoNotFound.Error()))

			return
		}
		c.Error(err)

		return
	}

	// delete photo data in repository permanently
	if err := pc.database.Unscoped().Model(&models.Photo{}).Delete(photo).Error; err != nil {
		c.Error(err)

		return
	}

	// send response with message
	c.JSON(http.StatusOK, gin.H{
		"message": "photo deleted",
	})
}
