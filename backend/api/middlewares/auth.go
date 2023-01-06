package middleware

import (
	"log"
	"net/http"
	authservice "ybigta/bard-backend/services/auth"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ok := authservice.AuthenticateSession(ctx)

		log.Default().Println("AuthenticateUser: ", ok)

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}
		ctx.Next()
	}
}
