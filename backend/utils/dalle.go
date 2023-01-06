package util

import (
	"log"
	"os"

	"github.com/astralservices/go-dalle"
)

type DALLEClient struct {
	c *dalle.Client
}

var DALLE *DALLEClient

func InitializeDALLE() {

	client := dalle.NewClient(os.Getenv("OPENAI_API_KEY"))

	DALLE = &DALLEClient{
		c: &client,
	}
}

type ImageGenerationReq struct {
	Summary string
}

func (d *DALLEClient) GenerateImage(req ImageGenerationReq) (string, error) {

	client := *d.c
	user := "BARD"
	response_type := "url"
	image_num := 1

	prompt := "A watercolor of a '" + req.Summary + "'."

	log.Default().Println("Prompt: ", prompt)

	resp, err := client.Generate(prompt, nil, &image_num, &user, &response_type)
	if err != nil {
		return "", err
	}

	log.Default().Println("Response: ", resp)

	return resp[0].URL, nil

}
