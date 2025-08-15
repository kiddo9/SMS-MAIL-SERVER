package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)


var returnResponse map[string]interface{}

var jwtSecretKey string

func init(){
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
}

// ValidateToken validates the JWT token and returns the parsed token or an error
func ValidateToken(Token string) (*jwt.Token, error) {
	// Parse the token using the secret key
	return jwt.Parse(Token, func(token*jwt.Token)(interface{}, error){
		// Check if the token's signing method is HMAC
		// This is a security check to ensure the token was signed with the expected method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			// If the signing method is not HMAC, return an error
			return nil, fmt.Errorf("error occoured")
		}
		// Return the secret key used to sign the token
		return []byte(jwtSecretKey), nil	
	})
}