package config

import "time"

const (
	ACCESS_TOKEN_COOKIE_NAME = "access_token"
	OAUTH_STATE_COOKIE_NAME  = "state"
	APP_IDENTIFIER           = "meline"
	ACCESS_TOKEN_EXPIRE      = time.Hour * 24 * 7
)
