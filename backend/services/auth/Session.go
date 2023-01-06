package authservice

import (
	"log"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateSession(ctx *gin.Context, userId uint) {

	session := sessions.Default(ctx)

	session.Set("Authenticated", true)
	session.Set("UserId", userId)
	session.Set("ExpiresAt", time.Now().Add(time.Hour*24*7).Unix())

	// credentials should be set on client side
	session.Save()
}

func DestroySession(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Save()
}

func AuthenticateSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)

	log.Default().Println("AuthenticateSession: ", session)

	if session.Get("Authenticated") == true {
		if session.Get("ExpiresAt").(int64) > time.Now().Unix() {
			return true
		}
	}

	return false
}

func GetSessionUserId(ctx *gin.Context) (uint, bool) {
	session := sessions.Default(ctx)
	userid := session.Get("UserId")
	if userid == nil {
		return 0, false
	}

	user_id := userid.(uint)

	return user_id, true
}
