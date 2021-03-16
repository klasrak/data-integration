package auth

import (
	di "github.com/klasrak/data-integration"
)

type AuthInterface interface {
	CreateAuth(string, *di.TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*di.AccessDetails) error
}
