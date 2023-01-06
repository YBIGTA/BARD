package storyservice

import (
	model "ybigta/bard-backend/models"
	repository "ybigta/bard-backend/repository/db"

	"github.com/gin-gonic/gin"
)

func GetStoriesByUserId(ctx *gin.Context, user_id uint) ([]model.Story, error) {

	stories := []model.Story{}

	err := repository.DB.Where("user_id = ?", user_id).Find(&stories).Error
	if err != nil {
		return nil, err
	}

	return stories, nil
}
