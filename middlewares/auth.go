package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/config"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/helpers"
	"gorm.io/gorm"
	"net/http"
)

type authMiddleware struct {
	serverConfig config.Config
	database     *gorm.DB
}

func NewAuthMiddleware(database *gorm.DB, serverConfig config.Config) *authMiddleware {
	am := authMiddleware{
		serverConfig: serverConfig,
		database:     database,
	}

	return &am
}

func (am *authMiddleware) Authorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": errAccessTokenRequired.Error(),
		})

		return
	}

	tokenClaims, err := helpers.ValidateJWT(accessToken, am.serverConfig.JwtSecretKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.Set("claims", tokenClaims.Claims)

	c.Next()
}
