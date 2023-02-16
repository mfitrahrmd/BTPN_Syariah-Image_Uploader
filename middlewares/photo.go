package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/config"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/helpers"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/models"
	"gorm.io/gorm"
	"net/http"
)

var (
	errAccessTokenRequired = errors.New("access token required")
	errAccessTokenInvalid  = errors.New("invalid access token")
	errInternalServer      = errors.New("unexpected server error, please try again later")
	errForbidden           = errors.New("you don't have permission to access this resource")
	errPhotoNotFound       = errors.New("photo not found")
)

type photoMiddleware struct {
	serverConfig config.Config
	database     *gorm.DB
}

func NewPhotoMiddleware(database *gorm.DB, serverConfig config.Config) *photoMiddleware {
	pm := photoMiddleware{
		serverConfig: serverConfig,
		database:     database,
	}

	return &pm
}

func (pm *photoMiddleware) VerifyPhotoOwner(c *gin.Context) {
	tokenClaims, ok := c.MustGet("claims").(helpers.Claims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	photoId := c.Param("photoId")

	var photo models.Photo

	if err := pm.database.Model(&models.Photo{}).First(&photo, photoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": errPhotoNotFound.Error(),
			})

			return
		}
	}

	if tokenClaims.UserID != photo.UserID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": errForbidden.Error(),
		})

		return
	}

	c.Next()
}
