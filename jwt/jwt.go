package jwt

import (
	"net/http"

	di "github.com/klasrak/data-integration"
)

type Jwt interface {
	CreateToken(di.User) (string, error)
	ExtractTokenMetadata(*http.Request)
}
