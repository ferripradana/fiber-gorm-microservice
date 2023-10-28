package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

const (
	Access  = "access"
	Refresh = "refresh"
)

type AppToken struct {
	Token          string    `json:"token"`
	TokenType      string    `json:"type"`
	ExpirationTime time.Time `json:"expiration_time"`
}

type Claims struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

var TokenTypeKeyName = map[string]string{
	Access:  "Secure.JWTAccessSecure",
	Refresh: "Secure.JWTRefreshSecure",
}

var TokenTypeExpTime = map[string]string{
	Access:  "Secure.JWTAccessTimeMinute",
	Refresh: "Secure.JWTRefreshTimeHour",
}

func GenerateJWTToken(userId int, tokenType string) (appToken *AppToken, err error) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("Fatal error in config file: %s", err.Error())
	}

	JWTSecureKey := viper.GetString(TokenTypeKeyName[tokenType])
	JWTExpTime := viper.GetString(TokenTypeExpTime[tokenType])

	tokenTimeConverted, err := strconv.ParseInt(JWTExpTime, 10, 64)
	if err != nil {
		return
	}

	tokenTimeUnix := time.Duration(tokenTimeConverted)
	switch tokenType {
	case Refresh:
		tokenTimeUnix *= time.Hour
	case Access:
		tokenTimeUnix *= time.Minute
	default:
		err = errors.New("Invalid token type")
	}
	if err != nil {
		return
	}

	nowTime := time.Now()
	expirationTokenTime := nowTime.Add(tokenTimeUnix)
	tokenClaims := &Claims{
		ID:   userId,
		Type: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTokenTime),
		},
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenStr, err := tokenWithClaims.SignedString([]byte(JWTSecureKey))
	if err != nil {
		return
	}

	appToken = &AppToken{
		Token:          tokenStr,
		TokenType:      tokenType,
		ExpirationTime: expirationTokenTime,
	}

	return
}
