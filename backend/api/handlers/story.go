package handler

import (
	"net/http"
	authservice "ybigta/bard-backend/services/auth"
	storyservice "ybigta/bard-backend/services/story"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context) {

	story, err := storyservice.CreateStory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Story created successfully", "story_id": story.ID})

}

func GetStoryByUserId(c *gin.Context) {

	user_id, ok := authservice.GetSessionUserId(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not logged in"})
		return
	}

	stories, err := storyservice.GetStoriesByUserId(c, user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stories": stories})

}

func GetStoryById(c *gin.Context) {

	story, err := storyservice.GetStory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"story": story})

}

func GetStoriesByUser(c *gin.Context) {

	user_id, ok := authservice.GetSessionUserId(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not logged in"})
		return
	}

	stories, err := storyservice.GetStoriesByUserId(c, user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stories": stories})

}

func UpdateStoryTitle(c *gin.Context) {

	story, err := storyservice.UpdateStoryTitle(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Story updated successfully", "story": story})
}
