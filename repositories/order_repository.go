package repositories

import (
	"depoguna-api/models"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	FindAll(paginate func(db *gorm.DB) *gorm.DB) (*[]models.Order, error)
	GetDetail(id int) (*models.Order, error)
	Insert(order *models.Order) error
	Update(req interface{}, id int) error
	Delete(id int) error
	Search(keyword int) (*[]models.Order, error)
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (r *orderRepository) FindAll(paginate func(db *gorm.DB) *gorm.DB) (*[]models.Order, error) {
	var orders []models.Order
	if err := r.DB.Scopes(paginate).Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

func (r *orderRepository) GetDetail(id int) (*models.Order, error) {
	var order models.Order
	if err := r.DB.Table("orders").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) Insert(order *models.Order) error {
	return r.DB.Table("orders").Create(&order).Error
}

func (r *orderRepository) Update(req interface{}, id int) error {
	var order models.Order
	if err := r.DB.Table("orders").Clauses(clause.OnConflict{DoNothing: true}).Model(order).Where("id = ?", id).Updates(req); err.Error != nil {
		return err.Error
	}
	r.DB.Model(&order).Where("id = ?", id).Update("UpdatedAt", time.Now())
	return nil
}

func (r *orderRepository) Delete(id int) error {
	var order models.Order
	if res := r.DB.Table("orders").Clauses(clause.Returning{}).Where("id = ?", id).Delete(&order); res.RowsAffected < 1 {
		return fmt.Errorf("order does not exist")
	}
	return nil
}

func (r *orderRepository) Search(keyword int) (*[]models.Order, error) {
	var order []models.Order
	if err := r.DB.Table("orders").Where("product_id = ?", keyword).Or("qty = ?", keyword).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
