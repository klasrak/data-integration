package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	di "github.com/klasrak/data-integration"
	"github.com/twinj/uuid"
)

type TokenManager struct{}

func NewTokenServer() *TokenManager {
	return &TokenManager{}
}

type TokenInterface interface {
	CreateToken(userID, email, jwtSecret string) (*di.TokenDetails, error)
	ExtractTokenMetadata(r *http.Request, jwtSecret string) (*di.AccessDetails, error)
}

var _ TokenInterface = &TokenManager{}

func (t *TokenManager) CreateToken(userID, email, jwtSecret string) (*di.TokenDetails, error) {
	td := &di.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 60).Unix() // expires after 1 hour
	td.TokenUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = fmt.Sprintf("%s++%s", td.TokenUUID, userID)

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["accessUUID"] = td.TokenUUID
	atClaims["userID"] = userID
	atClaims["email"] = email
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = at.SignedString([]byte(jwtSecret))

	if err != nil {
		return nil, err
	}

	// Create refresh token
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = fmt.Sprintf("%s++%s", td.TokenUUID, userID)

	rtClaims := jwt.MapClaims{}
	rtClaims["refreshUUID"] = td.RefreshUUID
	rtClaims["userID"] = userID
	rtClaims["email"] = email
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(jwtSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (t *TokenManager) ExtractTokenMetadata(r *http.Request, jwtSecret string) (*di.AccessDetails, error) {
	token, err := VerifyToken(r, jwtSecret)

	if err != nil {
		return nil, err
	}

	acc, err := Extract(token)

	if err != nil {
		return nil, err
	}

	return acc, nil
}

func TokenValid(r *http.Request, jwtSecret string) error {
	token, err := VerifyToken(r, jwtSecret)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func VerifyToken(r *http.Request, jwtSecret string) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// ExtractToken get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Extract(token *jwt.Token) (*di.AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUUID, ok := claims["accessUUID"].(string)
		userID, userOK := claims["userID"].(string)
		email, emailOk := claims["email"].(string)

		if ok == false || userOK == false || emailOk == false {
			return nil, errors.New("unauthorized")
		} else {
			return &di.AccessDetails{
				TokenUUID: accessUUID,
				UserID:    userID,
				Email:     email,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func ExtractTokenMetadata(r *http.Request, jwtSecret string) (*di.AccessDetails, error) {
	token, err := VerifyToken(r, jwtSecret)
	if err != nil {
		return nil, err
	}
	acc, err := Extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
