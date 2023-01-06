package storyservice

import (
	model "ybigta/bard-backend/models"
	repository "ybigta/bard-backend/repository/db"

	"github.com/gin-gonic/gin"
)

func GetStory(ctx *gin.Context) (*model.Story, error) {

	story_id := ctx.Params.ByName("story_id")

	story := model.Story{}

	err := repository.DB.Where("id = ?", story_id).First(&story).Error
	if err != nil {
		return nil, err
	}

	return &story, nil
}
