package storyservice

import (
	model "ybigta/bard-backend/models"
	repository "ybigta/bard-backend/repository/db"
	authservice "ybigta/bard-backend/services/auth"

	"github.com/gin-gonic/gin"
)

type UpdateStoryTitlePayload struct {
	Title string `json:"title"`
}

func UpdateStoryTitle(ctx *gin.Context) (*model.Story, error) {

	var payload UpdateStoryTitlePayload

	storyId := ctx.Param("story_id")

	err := ctx.BindJSON(&payload)
	if err != nil {
		return nil, err
	}

	user_id, _ := authservice.GetSessionUserId(ctx)

	updatedStory := model.Story{
		Title:  payload.Title,
		UserId: user_id,
	}

	err = repository.DB.Model(&model.Story{}).Where("id = ?", storyId).Update("title", payload.Title).Error
	if err != nil {
		return nil, err
	}

	return &updatedStory, nil

}
