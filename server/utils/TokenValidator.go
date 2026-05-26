package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var jwtSecretKey string = os.Getenv("JWT_SECRET_KEY")


// ValidateToken validates the JWT token and returns the parsed token or an error
func ValidateToken(Token string) (*jwt.Token, error) {
	// Parse the token using the secret key
	return jwt.Parse(Token, func(token*jwt.Token)(interface{}, error){
		// Check if the token's signing method is HMAC
		// This is a security check to ensure the token was signed with the expected method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			// If the signing method is not HMAC, return an error
			fmt.Println(ok)
			return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		// Return the secret key used to sign the token
		return []byte(jwtSecretKey), nil	
	})
}