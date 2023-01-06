package handler

import (
	"net/http"
	authservice "ybigta/bard-backend/services/auth"
	userservice "ybigta/bard-backend/services/user"

	"github.com/gin-gonic/gin"
)

func GoogleSignIn(c *gin.Context) {
	userservice.GoogleSignIn(c)
}

func Logout(ctx *gin.Context) {
	authservice.DestroySession(ctx)
	ctx.JSON(http.StatusNoContent, gin.H{"message": "logout success"})
}

func GetSession(ctx *gin.Context) {
	userId, ok := authservice.GetSessionUserId(ctx)

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"userId": userId})
}
