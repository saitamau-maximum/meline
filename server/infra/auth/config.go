package auth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func NewGithubOAuthConf() *oauth2.Config {
	newConfig := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	return newConfig
}


