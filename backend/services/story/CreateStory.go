package storyservice

import (
	model "ybigta/bard-backend/models"
	repository "ybigta/bard-backend/repository/db"
	authservice "ybigta/bard-backend/services/auth"
	util "ybigta/bard-backend/utils"

	"github.com/gin-gonic/gin"
)

type CreateStoryPayload struct {
	Characters []string `json:"characters"`
	ImageIds   []int    `json:"image_ids"`
}

func CreateStory(ctx *gin.Context) (*model.Story, error) {

	var payload CreateStoryPayload

	err := ctx.BindJSON(&payload)
	if err != nil {
		return nil, err
	}

	var captions []string
	for _, imageId := range payload.ImageIds {
		var image model.File
		err = repository.DB.Where("id = ?", imageId).First(&image).Error
		if err != nil {
			return nil, err
		}
		captions = append(captions, image.Caption)
	}

	story, err := util.GPT3.GenerateStory(util.StoryGenerationReq{
		Captions:   captions,
		Characters: payload.Characters,
	})
	if err != nil {
		return nil, err
	}

	summary, err := util.GPT3.SummarizeStory(util.StorySummarizationReq{
		Story: story,
	})

	if err != nil {
		return nil, err
	}

	imageUrl, err := util.DALLE.GenerateImage(util.ImageGenerationReq{
		Summary: summary,
	})
	if err != nil {
		return nil, err
	}

	user_id, _ := authservice.GetSessionUserId(ctx)

	createdStory := model.Story{
		Body:           story,
		SummarizedBody: summary,
		ImageUrl:       imageUrl,
		UserId:         user_id,
	}

	err = repository.DB.Omit("title").Create(&createdStory).Error
	if err != nil {
		return nil, err
	}

	return &createdStory, nil
}
