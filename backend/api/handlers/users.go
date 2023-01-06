package handler

import (
	userservice "ybigta/bard-backend/services/user"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	userservice.CreateUser(ctx)
}
