package userservice

import (
	"net/http"
	model "ybigta/bard-backend/models"
	repository "ybigta/bard-backend/repository/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateUserPayload struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	SocialId string `json:"social_id"`
}

func CreateUser(ctx *gin.Context) {

	// check if socialId is already registered
	var createUserPayload CreateUserPayload

	err := ctx.BindJSON(&createUserPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if not registered -> bad req
	var social model.Social
	err = repository.DB.Where("social_id = ?", createUserPayload.SocialId).First(&social).Error
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Social ID is not registered"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if registered -> create user
	user := model.User{
		Email: createUserPayload.Email,
		Name:  createUserPayload.Name,
	}
	err = repository.DB.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// and update to social table

	err = repository.DB.Model(&social).Update("user_id", user.ID).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
