package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUtil interface {
	GenerateToken(userId int) string
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type jwtUtil struct {
	secretKey string
}

func NewJWTUtil() JWTUtil {
	return &jwtUtil{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey != "" {
		secretKey = "secret_key"
	}
	return secretKey
}

// GenerateToken creates a new token for a specific userId
func (j *jwtUtil) GenerateToken(userId int) string {
	claims := &jwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// VerifyToken checks if the token is valid or not
func (j *jwtUtil) VerifyToken(tokenString string) (*jwt.Token, error) {

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	}

	// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
