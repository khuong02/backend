package my_jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type JWTType int

const (
	Authentication JWTType = 0
)

type MyJWT struct {
	UserID   uuid.UUID
	UserName string
	JWTType  JWTType
	jwt.StandardClaims
}

func GenerateJWT(data MyJWT, secretKey string) (string, error) {
	claims := &MyJWT{
		UserID:   data.UserID,
		UserName: data.UserName,
		JWTType:  data.JWTType,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			IssuedAt:  data.IssuedAt,
			ExpiresAt: data.ExpiresAt,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", GenerateJWTFailed(err)
	}

	return tokenString, nil
}

func VerifyJWT(tokenStr string, secretKey string) (*MyJWT, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyJWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, JWTExpired(err)
		}

		if validationErr.Errors != jwt.ValidationErrorExpired {
			return nil, VerifyJWTUnauthorized(errors.New(validationErr.Error()))
		}

		return nil, JWTExpired(err)
	}

	if claims, ok := token.Claims.(*MyJWT); ok {
		return claims, nil
	}

	return nil, VerifyJWTFailed(errors.New("Invalid Claims"))
}
