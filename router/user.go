package router

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	POSTRegisterUser(c *gin.Context)
	GETLoginUser(c *gin.Context)
	PUTUpdateUser(c *gin.Context)
	DELETEDeleteUser(c *gin.Context)
}

func WithUserRoutes(r gin.IRouter, userController UserController) {
	users := r.Group("/users")
	{
		users.POST("/register", userController.POSTRegisterUser)
		users.GET("/login", userController.GETLoginUser)
		users.PUT("/:userId", userController.PUTUpdateUser)
		users.DELETE("/:userId", userController.DELETEDeleteUser)
	}
}
