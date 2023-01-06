package repository

import (
	"os"
	config "ybigta/bard-backend/configs"
	repository "ybigta/bard-backend/repository/db"

	"github.com/gin-contrib/sessions"
	gorm_sessions "github.com/gin-contrib/sessions/gorm"
)

var Sessionstore *gorm_sessions.Store

func ConnectSessionStore() {

	store := gorm_sessions.NewStore(repository.DB, config.EXPIRED_SESSION_CLEANUP, []byte(os.Getenv("BARDAUTH_SECRET")))
	store.Options(sessions.Options{
		MaxAge:   config.SESSION_MAX_AGE,
		HttpOnly: true,
		Path:     "/",
		// Secure:   true,
		// SameSite: http.SameSiteDefaultMode,
	})

	Sessionstore = &store

}
