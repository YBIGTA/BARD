package util

import (
	"context"
	"os"

	"google.golang.org/api/idtoken"
)

type GoogleUserInfo struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
}

func ValidateGoogleIdToken(tokenString string) (*GoogleUserInfo, error) {

	payload, err := idtoken.Validate(context.Background(), tokenString, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return nil, err
	}

	var tokenInfo GoogleUserInfo

	tokenInfo.Email = payload.Claims["email"].(string)
	tokenInfo.EmailVerified = payload.Claims["email_verified"].(bool)
	tokenInfo.Name = payload.Claims["name"].(string)
	tokenInfo.Picture = payload.Claims["picture"].(string)
	tokenInfo.Sub = payload.Claims["sub"].(string)

	return &tokenInfo, nil
}
