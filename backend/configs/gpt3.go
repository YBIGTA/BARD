package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type StoryGenerationPromptConfig struct {
	Temp      float32 `yaml:"temp"`
	MaxTokens int     `yaml:"max"`
	Task      string  `yaml:"task"`
	Example   struct {
		Characters []string `yaml:"characters"`
		Captions   []string `yaml:"captions"`
		Story      string   `yaml:"story"`
	}
}

type StorySummarizationPromptConfig struct {
	Temp      float32 `yaml:"temp"`
	MaxTokens int     `yaml:"max"`
	Task      string  `yaml:"task"`
}

type GPT3Config struct {
	StoryGenerationConfig    StoryGenerationPromptConfig    `yaml:"story_generation"`
	StorySummarizationConfig StorySummarizationPromptConfig `yaml:"story_summarization"`
}

var StoryGenerationConfig *StoryGenerationPromptConfig
var StorySummarizationConfig *StorySummarizationPromptConfig

func LoadGPT3Config() {
	var config *GPT3Config

	configPath := os.Getenv("GPT3_CONFIG_PATH")
	log.Default().Println("Loading GPT3 config from", configPath)

	yamlFile, err := os.ReadFile(os.Getenv("GPT3_CONFIG_PATH"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	StoryGenerationConfig = &config.StoryGenerationConfig
	StorySummarizationConfig = &config.StorySummarizationConfig

}
