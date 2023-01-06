package api

import (
	handler "ybigta/bard-backend/api/handlers"
	middleware "ybigta/bard-backend/api/middlewares"
	config "ybigta/bard-backend/configs"
	session_repository "ybigta/bard-backend/repository/session"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.Use(cors.New(config.CorsConfig()))
	router.Use(gin.Logger())

	router.Use(sessions.Sessions("bard-session", *session_repository.Sessionstore))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)

		v1.POST("/users", handler.CreateUser)
		auth := v1.Group("/auth")
		{
			auth.POST("/google", handler.GoogleSignIn)
			auth.GET("/user", middleware.AuthenticateUser(), handler.GetSession)
			auth.GET("/logout", middleware.AuthenticateUser(), handler.Logout)
		}

		v1.GET("/stories", middleware.AuthenticateUser(), handler.GetStoriesByUser)
		v1.POST("/stories", middleware.AuthenticateUser(), handler.CreateStory)

		stories_api := v1.Group("/stories", middleware.AuthenticateUser())
		{
			stories_api.GET("/:story_id", handler.GetStoryById)
			stories_api.PATCH("/:story_id/title", handler.UpdateStoryTitle)
		}

		// TODO : Add Authentication
		v1.POST("/files/upload", middleware.AuthenticateUser(), handler.UploadFile)

	}

	return router
}
