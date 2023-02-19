package middlewares

import (
	"github.com/gin-gonic/gin"
	controllerError "github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/error"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message,omitempty"`
	Details any    `json:"details,omitempty"`
}

func ErrorHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if isErrExist := len(c.Errors) > 0; isErrExist {
			for _, e := range c.Errors {

				if ce, ok := e.Err.(controllerError.ControllerError); ok {
					c.AbortWithStatusJSON(ce.StatusCode, errorResponse{
						Message: ce.Message,
						Details: ce.Details,
					})

					return
				}

				logger.Errorln(e)

				c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{
					Message: e.Error(),
				})

				return

			}
		}
	}
}
