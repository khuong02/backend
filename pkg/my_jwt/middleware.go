package my_jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/pkg/helper"
	"github.com/pkg/errors"
	"strings"
)

func handleClaims(c *gin.Context, key string) (interface{}, error) {
	val, ok := c.Get(key)
	if !ok {
		return nil, VerifyJWTFailed(errors.New("Invalid Params Claims"))
	}

	return val, nil
}

func VerifyJWTMiddleware(cfg config.Config) func(*gin.Context) {
	return func(c *gin.Context) {

		authHeader := c.Request.Header["Token"]
		if len(authHeader) == 0 {
			err := errors.New("Not found token")
			helper.Error(c, JWTExpired(err))

			return
		}

		token := strings.TrimPrefix(authHeader[0], "Bearer ")

		claims, err := VerifyJWT(token, cfg.JWT.AccessSecretKey)
		if err != nil {
			helper.Error(c, err.(helper.Err))

			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.UserName)
		c.Set("jwt_type", claims.JWTType)

		c.Next()
	}

}

func AuthenticationJWTType(c *gin.Context) {
	val, err := handleClaims(c, "jwt_type")
	if err != nil {
		helper.Error(c, err.(helper.Err))

		return
	}

	if val.(JWTType) != Authentication {
		helper.Error(
			c,
			VerifyJWTUnauthorized(errors.New("Unauthorized")),
		)

		return
	}

	c.Next()
}
