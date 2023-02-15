package router

import (
	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	POSTInsertPhoto(c *gin.Context)
	GETFindAllPhotos(c *gin.Context)
	PUTUpdatePhoto(c *gin.Context)
	DELETEDeletePhoto(c *gin.Context)
}

func WithPhotoRoutes(r gin.IRouter, photoController PhotoController) {
	users := r.Group("/photos")
	{
		users.POST("/", photoController.POSTInsertPhoto)
		users.GET("/", photoController.GETFindAllPhotos)
		users.PUT("/:photoId", photoController.PUTUpdatePhoto)
		users.DELETE("/:photoId", photoController.DELETEDeletePhoto)
	}
}
