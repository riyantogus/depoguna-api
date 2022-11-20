package middleware

import (
	"depoguna-api/helpers"
	"depoguna-api/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func JWTAuth(ctx *gin.Context) {

	authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Errors:  err.Error(),
		})
		ctx.Abort()
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Errors:  err.Error(),
		})
		ctx.Abort()
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Errors:  err.Error(),
		})
		ctx.Abort()
		return
	}

	accessToken := fields[1]
	jwtToken := utils.NewJWTUtil()
	_, err := jwtToken.VerifyToken(accessToken)
	if err != nil {
		err := errors.New("invalid token")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Errors:  err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}
