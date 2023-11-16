package config

import (
	"fmt"
	"os"
)

// REQUIRED ENV
const (
	_MYSQL_ROOT_PASSWORD     = "MYSQL_ROOT_PASSWORD"
	_MYSQL_DATABASE          = "MYSQL_DATABASE"
	_MYSQL_USER              = "MYSQL_USER"
	_MYSQL_PASSWORD          = "MYSQL_PASSWORD"
	_GITHUB_CLIENT_ID        = "GITHUB_CLIENT_ID"
	_GITHUB_CLIENT_SECRET    = "GITHUB_CLIENT_SECRET"
	_ENV                     = "ENV"
	_FRONT_OAUTH_SUCCESS_URL = "FRONT_OAUTH_SUCCESS_URL"
	_JWT_SECRET              = "JWT_SECRET"
)

var (
	MYSQL_ROOT_PASSWORD     = os.Getenv(_MYSQL_ROOT_PASSWORD)
	MYSQL_DATABASE          = os.Getenv(_MYSQL_DATABASE)
	MYSQL_USER              = os.Getenv(_MYSQL_USER)
	MYSQL_PASSWORD          = os.Getenv(_MYSQL_PASSWORD)
	GITHUB_CLIENT_ID        = os.Getenv(_GITHUB_CLIENT_ID)
	GITHUB_CLIENT_SECRET    = os.Getenv(_GITHUB_CLIENT_SECRET)
	ENV                     = os.Getenv(_ENV)
	FRONT_OAUTH_SUCCESS_URL = os.Getenv(_FRONT_OAUTH_SUCCESS_URL)
	JWT_SECRET              = os.Getenv(_JWT_SECRET)

	IsDev = ENV == "development"

	requiredDBEnv = []string{
		_MYSQL_ROOT_PASSWORD,
		_MYSQL_DATABASE,
		_MYSQL_USER,
		_MYSQL_PASSWORD,
	}
	requiredAppEnv = []string{
		_MYSQL_ROOT_PASSWORD,
		_MYSQL_DATABASE,
		_MYSQL_USER,
		_MYSQL_PASSWORD,
		_GITHUB_CLIENT_ID,
		_GITHUB_CLIENT_SECRET,
		_ENV,
		_FRONT_OAUTH_SUCCESS_URL,
		_JWT_SECRET,
	}
)

func ValidateDBEnv() error {
	for _, env := range requiredDBEnv {
		if os.Getenv(env) == "" {
			return fmt.Errorf("%s is not set", env)
		}
	}

	return nil
}

func ValidateAppEnv() error {
	for _, env := range requiredAppEnv {
		if os.Getenv(env) == "" {
			return fmt.Errorf("%s is not set", env)
		}
	}

	return nil
}

func GetEnv(name string, defaultValue string) string {
	if os.Getenv(name) == "" {
		return defaultValue
	}

	return os.Getenv(name)
}

func GetEnvStrict(name string, err error) {
	if os.Getenv(name) == "" {
		panic(err)
	}
}
