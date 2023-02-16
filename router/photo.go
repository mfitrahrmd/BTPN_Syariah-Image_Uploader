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

type PhotoMiddleware interface {
	VerifyPhotoOwner(c *gin.Context)
}

type AuthMiddleware interface {
	Authorization(c *gin.Context)
}

func WithPhotoRoutes(r gin.IRouter, photoController PhotoController, authMiddleware AuthMiddleware, photoMiddleware PhotoMiddleware) {
	photos := r.Group("/photos")
	{
		photos.Use(authMiddleware.Authorization)
		photos.POST("", photoController.POSTInsertPhoto)
		photos.GET("", photoController.GETFindAllPhotos)
		photos.Use(photoMiddleware.VerifyPhotoOwner)
		photos.PUT("/:photoId", photoController.PUTUpdatePhoto)
		photos.DELETE("/:photoId", photoController.DELETEDeletePhoto)
	}
}
