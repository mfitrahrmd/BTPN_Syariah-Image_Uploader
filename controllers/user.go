package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/app"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/config"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/helpers"
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var (
	errInternalServer       = errors.New("unexpected server error, please try again later")
	errUserNotFound         = errors.New("user not found")
	errWrongPassword        = errors.New("wrong password")
	errUsernameAlreadyExist = errors.New("username already exist")
	errEmailAlreadyExist    = errors.New("email already exist")
)

type userController struct {
	serverConfig config.Config
	database     *gorm.DB
}

// NewUserController create instance of user controller
func NewUserController(database *gorm.DB, serverConfig config.Config) *userController {
	uc := userController{
		serverConfig: serverConfig,
		database:     database,
	}

	return &uc
}

func (uc *userController) POSTRegisterUser(c *gin.Context) {
	// bind and check user request json data
	var req app.RegisterUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": helpers.CustomValidationErrorMessage(err.(validator.ValidationErrors)),
		})

		return
	}

	// check if username already in used
	q := uc.database.Model(&models.User{}).Where("username = ?", req.Username).Find(&models.User{})
	if q.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	if q.RowsAffected > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errUsernameAlreadyExist.Error(),
		})

		return
	}

	// check if email already in used
	q = uc.database.Model(&models.User{}).Where("email = ?", req.Email).Find(&models.User{})
	if q.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	if q.RowsAffected > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errEmailAlreadyExist.Error(),
		})

		return
	}

	// hash user's password from request
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer,
		})

		return
	}

	// save user data into database repository
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err = uc.database.Model(&models.User{}).Create(&user).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer,
		})

		return
	}

	// send response with saved user's data
	c.JSON(http.StatusCreated, app.RegisterUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}

func (uc *userController) GETLoginUser(c *gin.Context) {
	// bind and check user request json data
	var req app.LoginUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// check if user is exists in repository with given email
	var user models.User

	if err := uc.database.Model(&models.User{}).First(&user, "email = ?", req.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": errUserNotFound.Error(),
			})

			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	// check if given password is correct
	isMatch, err := helpers.ComparePassword(req.Password, user.Password)
	if err != nil {
		log.Println(errors.Is(err, bcrypt.ErrMismatchedHashAndPassword))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	if !isMatch {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": errWrongPassword.Error(),
		})

		return
	}

	// generate access token
	token, err := helpers.GenerateJWT(helpers.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(uc.serverConfig.JwtTokenExpirationLength).Unix(),
		},
		Claims: helpers.Claims{
			UserID: user.ID,
		},
	}, uc.serverConfig.JwtSecretKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	// send response with generated access token
	c.JSON(http.StatusOK, app.LoginUserResponse{
		AccessToken: token,
	})
}

func (uc *userController) PUTUpdateUser(c *gin.Context) {
	userId := c.Param("userId")

	// bind and check user request json data
	var req app.UpdateUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// check if user to be updated is exists in repository
	var user models.User

	if err := uc.database.Model(&models.User{}).First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": errUserNotFound.Error(),
			})

			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return

	}

	// check if username already in used
	q := uc.database.Model(&models.User{}).Where("username = ?", req.Username).Not("id = ?", userId).Find(&models.User{})
	if q.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	if q.RowsAffected > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errUsernameAlreadyExist.Error(),
		})

		return
	}

	// check if email already in used
	q = uc.database.Model(&models.User{}).Where("email = ?", req.Email).Not("id = ?", userId).Find(&models.User{})
	if q.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	if q.RowsAffected > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errEmailAlreadyExist.Error(),
		})

		return
	}

	// update the user data
	user.Username = req.Username
	user.Email = req.Email

	if err = uc.database.Model(&models.User{}).Where("id = ?", userId).Updates(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	// send response with updated user's data
	c.JSON(http.StatusOK, app.UpdateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}

func (uc *userController) DELETEDeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	// check if user to be deleted is exists in repository
	var user models.User

	if err := uc.database.Model(&models.User{}).First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": errUserNotFound.Error(),
			})

			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return

	}

	// delete user data in repository permanently
	if err := uc.database.Unscoped().Model(&models.User{}).Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errInternalServer.Error(),
		})

		return
	}

	// send response with message
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}
