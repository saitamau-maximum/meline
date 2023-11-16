package github

import (
	"github.com/saitamau-maximum/meline/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func NewGithubOAuthConf() *oauth2.Config {
	newConfig := &oauth2.Config{
		ClientID:     config.GITHUB_CLIENT_ID,
		ClientSecret: config.GITHUB_CLIENT_SECRET,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	return newConfig
}
