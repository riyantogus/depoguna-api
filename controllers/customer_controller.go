package controllers

import (
	"depoguna-api/dto"
	"depoguna-api/helpers"
	"depoguna-api/models"
	"depoguna-api/repositories"
	"depoguna-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type CustomerController interface {
	FindAll(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Search(ctx *gin.Context)
}

type customerController struct {
	Repository repositories.CustomerRepository
	Validation utils.ValidationUtil
}

func NewCustomerController(db *gorm.DB) CustomerController {
	return &customerController{
		Repository: repositories.NewCustomerRepository(db),
		Validation: utils.NewValidationUtil(),
	}
}

func paginate(ctx *gin.Context) (offset, pageSize int) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset = (page - 1) * pageSize

	return offset, pageSize
}

func (c *customerController) FindAll(ctx *gin.Context) {
	offset, pageSize := paginate(ctx)
	customers, err := c.Repository.FindAll(offset, pageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	var res []dto.CustomerResponse
	for _, each := range *customers {
		res = append(res, dto.CustomerResponse(each))
	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (c *customerController) GetDetail(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	customer, err := c.Repository.GetDetail(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	res := dto.CustomerResponse(*customer)

	ctx.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (c *customerController) Insert(ctx *gin.Context) {
	var req dto.CustomerInput
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	if err := c.Validation.Validate(req); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "something went wrong. please try again",
				Errors:  c.Validation.ErrorMessage(fieldErr),
			})
			return
		}
	}

	customer := models.Customer{
		UserId:      req.UserId,
		Name:        req.Name,
		Email:       strings.ToLower(req.Email),
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		Mobile:      req.Mobile,
		Address:     req.Address,
	}

	if err := c.Repository.Insert(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	res := dto.CustomerResponse(customer)

	ctx.JSON(http.StatusCreated, helpers.Response{
		Status:  http.StatusCreated,
		Message: "customer has been created",
		Data:    res,
	})
}

func (c *customerController) Update(ctx *gin.Context) {
	var req dto.CustomerInput
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	if err := c.Validation.Validate(req); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "something went wrong. please try again",
				Errors:  c.Validation.ErrorMessage(fieldErr),
			})
			return
		}
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.Repository.Update(req, id); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, helpers.DefaultResponse{
		Status:  http.StatusOK,
		Message: "customer detail have been updated",
	})
}

func (c *customerController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.Repository.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, helpers.DefaultResponse{
		Status:  http.StatusOK,
		Message: "customer has been deleted",
	})
}

func (c *customerController) Search(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	customers, err := c.Repository.Search(keyword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	var res []dto.CustomerResponse
	for _, each := range *customers {
		res = append(res, dto.CustomerResponse(each))
	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}
