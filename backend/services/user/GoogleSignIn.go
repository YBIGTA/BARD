package userservice

import (
	"net/http"
	model "ybigta/bard-backend/models"
	repository "ybigta/bard-backend/repository/db"
	authservice "ybigta/bard-backend/services/auth"
	util "ybigta/bard-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GoogleAuth struct {
	Credential string `json:"credential"`
	ClientId   string `json:"clientId"`
}

func GoogleSignIn(ctx *gin.Context) {

	var googleAuthPayload GoogleAuth

	err := ctx.BindJSON(&googleAuthPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate google id token
	token_info, err := util.ValidateGoogleIdToken(googleAuthPayload.Credential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if social ID is already registered
	var social model.Social
	err = repository.DB.Where("social_id = ?", token_info.Sub).First(&social).Error

	// if social_id is not registered, Create Social ID and request for signup
	if err == gorm.ErrRecordNotFound {
		social_err := repository.DB.Omit("user_id").Create(model.Social{Id: token_info.Sub, Email: token_info.Email, Provider: "google"}).Error
		if social_err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Not registered. Please sign up first", "social_id": token_info.Sub, "email": token_info.Email, "name": token_info.Name, "provider": "google"})
		return
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if social_id is already registered, check user exists
	var user model.User
	err = repository.DB.Where("id = ?", social.UserId).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusOK, gin.H{"message": "Not registered. Please sign up first", "social_id": token_info.Sub, "email": token_info.Email, "name": token_info.Name, "provider": "google"})
		return
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authservice.CreateSession(ctx, user.ID)
	ctx.JSON(http.StatusCreated, gin.H{"message": "User logged in successfully", "userID": user.ID})

}
