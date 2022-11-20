package dto

import "time"

type OrderInput struct {
	CustomerId int `json:"customer_id" form:"customer_id" validate:"required"`
	ProductId  int `json:"product_id" form:"product_id" validate:"required"`
	Qty        int `json:"qty" form:"qty" validate:"required"`
}

type OrderResponse struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

type OrderDetailResponse struct {
	Id         int       `json:"id"`
	CustomerId int       `json:"customer_id"`
	ProductId  int       `json:"product_id"`
	Qty        int       `json:"qty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
