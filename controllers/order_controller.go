package controllers

import (
	"depoguna-api/dto"
	"depoguna-api/helpers"
	"depoguna-api/models"
	"depoguna-api/repositories"
	"depoguna-api/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type OrderController interface {
	FindAll(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Search(ctx *gin.Context)
}

type orderController struct {
	Repository repositories.OrderRepository
	Validation utils.ValidationUtil
}

func NewOrderController(db *gorm.DB) OrderController {
	return &orderController{
		Repository: repositories.NewOrderRepository(db),
		Validation: utils.NewValidationUtil(),
	}
}

func (c *orderController) FindAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	paginate := utils.Paginate(page, pageSize)
	orders, err := c.Repository.FindAll(paginate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	var res []dto.OrderResponse
	for _, each := range *orders {
		res = append(res, dto.OrderResponse{
			Id:        each.Id,
			ProductId: each.ProductId,
			Qty:       each.Qty,
		})
	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (c *orderController) GetDetail(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	order, err := c.Repository.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, helpers.DefaultResponse{
				Status:  http.StatusNotFound,
				Message: "order does not exist",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	res := dto.OrderDetailResponse(*order)

	ctx.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (c *orderController) Insert(ctx *gin.Context) {
	var req dto.OrderInput
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

	order := models.Order{
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
		Qty:        req.Qty,
	}

	if err := c.Repository.Insert(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	res := dto.OrderDetailResponse(order)

	ctx.JSON(http.StatusCreated, helpers.Response{
		Status:  http.StatusCreated,
		Message: "order has been created",
		Data:    res,
	})
}

func (c *orderController) Update(ctx *gin.Context) {
	var req dto.OrderInput
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
		Message: "order detail have been updated",
	})
}

func (c *orderController) Delete(ctx *gin.Context) {
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
		Message: "order has been deleted",
	})
}

func (c *orderController) Search(ctx *gin.Context) {
	keyword, _ := strconv.Atoi(ctx.Query("keyword"))
	orders, err := c.Repository.Search(keyword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "something went wrong. please try again",
			Errors:  err.Error(),
		})
		return
	}

	var res []dto.OrderResponse
	for _, each := range *orders {
		res = append(res, dto.OrderResponse{
			Id:        each.Id,
			ProductId: each.ProductId,
			Qty:       each.Qty,
		})
	}

	ctx.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}
