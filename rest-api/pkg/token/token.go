package token

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context của JSON web token
type Context struct {
	ID       uint64
	Username string
}

// Validates the secret format
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// Parse validates the token with secret
func Parse(tokenString string, secret string, c *gin.Context) (*Context, error) {
	ctx := &Context{}

	// Parse token
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		c.Set("userID", ctx.ID)
		return ctx, nil
	} else {
		return ctx, err
	}
}

// ParseRequest
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string

	// Parse header to get token part
	fmt.Sscanf(header, "%s", &t)
	return Parse(t, secret, c)
}

// Sign - Register context with secret
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err = token.SignedString([]byte(secret))
	return
}
