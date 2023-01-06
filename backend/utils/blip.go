package util

import (
	"context"
	"os"
	"ybigta/bard-backend/pkg/replicate"
)

var BLIP *replicate.Model

type Task string

const (
	ImageCaptioning   Task = "image_captioning"
	QuestionAnswering Task = "visual_question_answering"
	TextMatching      Task = "image_text_matching"
)

type BLIPInput struct {
	// url of the input image
	Image string `json:"image"`
	Task  string `json:"task"`
	// Type question for the input image for visual question answering task.
	Question string `json:"question"`
	// Type caption for the input image for image text matching task.
	Caption string `json:"caption"`
}

type BlipPayload struct {
	Inputs BLIPInput `json:"inputs"`
}

func InitializeBLIP() {

	client := replicate.NewClient(os.Getenv("REPLICATE_API_KEY"))

	context := context.Background()

	modelInfo, err := client.GetModel(context, replicate.GetModelReqParams{
		Owner: "salesforce",
		Name:  "blip",
	})
	if err != nil {
		panic(err)
	}

	model := client.NewModel(replicate.ModelConfig{
		Version:          modelInfo.LatestVersion.Id,
		WebhookCompleted: "",
	})

	BLIP = &model
}
