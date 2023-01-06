package util

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	config "ybigta/bard-backend/configs"

	"github.com/PullRequestInc/go-gpt3"
)

type GPT3Client struct {
	c *gpt3.Client
}

type StoryGenerationReq struct {
	Captions   []string
	Characters []string
}

type StorySummarizationReq struct {
	Story string
}

var GPT3 *GPT3Client

func InitializeGPT3() {

	client := gpt3.NewClient(os.Getenv("OPENAI_API_KEY"), gpt3.WithDefaultEngine(gpt3.TextDavinci003Engine), gpt3.WithTimeout(time.Duration(120*time.Second)))

	GPT3 = &GPT3Client{
		c: &client,
	}
}

func (g *GPT3Client) GenerateStory(req StoryGenerationReq) (string, error) {

	example_character_prompt := generateCharactersPrompt(config.StoryGenerationConfig.Example.Characters)
	example_captions_prompt := generateCaptionsPrompt(config.StoryGenerationConfig.Example.Captions)
	example_story := config.StoryGenerationConfig.Example.Story

	character_prompt := generateCharactersPrompt(req.Characters)
	captions_prompt := generateCaptionsPrompt(req.Captions)

	task := config.StoryGenerationConfig.Task
	max_tokens := config.StoryGenerationConfig.MaxTokens
	temp := config.StoryGenerationConfig.Temp

	prompt :=
		task + " Example:\n\n" +
			example_character_prompt + "\n" +
			example_captions_prompt + "\n" +
			"Story: " + example_story + "\n\n" +
			character_prompt + "\n" +
			captions_prompt + "\n" +
			"Story: "

	client := *g.c
	ctx := context.Background()

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:      []string{prompt},
		MaxTokens:   gpt3.IntPtr(max_tokens),
		Temperature: &temp,
	})
	if err != nil {
		log.Default().Println(err)
		return "", err
	}

	story := resp.Choices[0].Text

	return story, nil
}

func (g *GPT3Client) SummarizeStory(req StorySummarizationReq) (string, error) {

	task := config.StorySummarizationConfig.Task
	max_tokens := config.StorySummarizationConfig.MaxTokens
	temp := config.StorySummarizationConfig.Temp

	prompt := "Story: " + req.Story + "\n\n" +
		"Task: " + task + "\n\n" +
		"Summarization: This is a story about"

	client := *g.c
	ctx := context.Background()

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:      []string{prompt},
		MaxTokens:   gpt3.IntPtr(max_tokens),
		Temperature: &temp,
	})
	if err != nil {
		log.Default().Println(err)
		return "", err
	}

	summary := strings.Trim(resp.Choices[0].Text, " ")

	return summary, nil

}

func generateCharactersPrompt(characters []string) string {
	prompt := "Characters: "

	prompt += strings.Join(characters, ", ")

	return prompt
}

func generateCaptionsPrompt(captions []string) string {

	for i := range captions {
		captions[i] += "."
	}

	first_caption := captions[0]
	last_caption := captions[len(captions)-1]
	middle_captions := captions[1 : len(captions)-1]

	prompt := fmt.Sprintf("Events: %s %s %s", "First, "+first_caption, "Then, "+strings.Join(middle_captions, " Then, "), "Finally, "+last_caption)

	return prompt
}
