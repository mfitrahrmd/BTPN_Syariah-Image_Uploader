package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/controllers"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/helpers"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ErrorHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if isErrExist := len(c.Errors) > 0; isErrExist {
			for _, e := range c.Errors {
				if e.IsType(gin.ErrorTypeBind) {
					if vErr, ok := e.Err.(validator.ValidationErrors); ok {
						c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"message": helpers.CustomValidationErrorMessage(vErr),
						})

						return
					}
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "invalid request data",
					})

					return
				}

				if e.IsType(gin.ErrorTypePublic) {
					var cErr controllers.ControllerError

					if errors.As(e.Err, &cErr) {
						c.AbortWithStatusJSON(cErr.StatusCode, gin.H{
							"message": cErr.Error(),
						})

						return
					}

					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": e.Error(),
					})

					return
				}

				{
					logger.Errorln(e)

					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"message": errInternalServer.Error(),
					})

					return
				}

			}
		}
	}
}
