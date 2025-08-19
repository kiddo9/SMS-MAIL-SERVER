package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string = os.Getenv("JWT_SECRET_KEY")

func GenerateJWTToken(email string, uuid string, APIKey string, Time int64) (string, error) {
	//create claims (that is data that will be encoded in the JWT)
	claims := jwt.MapClaims{
		"email":   email,
		"uuid":    uuid,
		"APIKey":  APIKey,
		"otp_expires_at": Time,
		"exp": time.Now().Add(time.Minute * 10).Unix(), // token expires after 3 minutes
		"iat": time.Now().Unix(), // issued at time
	}

	//create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
func GenerateJWTTokenLongTerm(email string, uuid string, APIKey string) (string, error) {
	//create claims (that is data that will be encoded in the JWT)
	claims := jwt.MapClaims{
		"email":   email,
		"uuid":    uuid,
		"APIKey":  APIKey,
		"exp": time.Now().Add(time.Hour *8760).Unix(), // token expires after 1 year
		"iat": time.Now().Unix(), // issued at time
	}

	//create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func GenerateRequestJWTToken(uuid string, APIKey string) (string, error) {
	//create claims (that is data that will be encoded in the JWT)
	claims := jwt.MapClaims{
		"uuid":    uuid,
		"APIKey":  APIKey,
		"exp": time.Now().Add(time.Minute * 5).Unix(), // token expires after 5 minutes
		"iat": time.Now().Unix(), // issued at time
	}

	//create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}