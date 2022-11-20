package controllers

import (
	"depoguna-api/dto"
	"depoguna-api/helpers"
	"depoguna-api/models"
	"depoguna-api/repositories"
	"depoguna-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	DB         *gorm.DB
	Repository repositories.AuthRepository
	Validation utils.ValidationUtil
	JWT        utils.JWTUtil
	Encrypt    utils.EncryptUtil
}

func NewAuthController(db *gorm.DB) AuthController {
	return &authController{
		DB:         db,
		Repository: repositories.NewAuthRepository(db),
		Validation: utils.NewValidationUtil(),
		JWT:        utils.NewJWTUtil(),
		Encrypt:    utils.NewEncryptUtil(),
	}
}

func (a *authController) Register(ctx *gin.Context) {
	var req dto.RegisterInput
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	if err := a.Validation.Validate(req); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "something went wrong. please try again",
				Errors:  a.Validation.ErrorMessage(fieldErr),
			})
			return
		}
	}

	user := models.User{
		Name:     req.Name,
		Email:    strings.ToLower(req.Email),
		Password: a.Encrypt.HashAndSalt([]byte(req.Password)),
	}

	if a.Repository.IsDuplicateEmail(user.Email) {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  "email already exist",
		})
		return
	}

	if err := a.Repository.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	res := dto.RegisterResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, helpers.Response{
		Status:  http.StatusCreated,
		Message: "thanks for signing up. your account has been created",
		Data:    res,
	})
}

func (a *authController) Login(ctx *gin.Context) {
	var req dto.LoginInput
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	if err := a.Validation.Validate(req); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Something went wrong. Please try again.",
				Errors:  a.Validation.ErrorMessage(fieldErr),
			})
			return
		}
	}

	user, err := a.Repository.Login(strings.ToLower(req.Email))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "something went wrong. please try again",
			Errors:  "invalid email or password",
		})
		return
	}

	if checkPwd := a.Encrypt.VerifyPassword(user.Password, []byte(req.Password)); !checkPwd {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  "invalid email or password",
		})
		return
	}

	jwtToken := a.JWT.GenerateToken(user.Id)

	res := dto.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: jwtToken,
	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "you are successfully logged in",
		Data:    res,
	})
}
