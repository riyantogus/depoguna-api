package repositories

import (
	"depoguna-api/models"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomerRepository interface {
	FindAll(paginate func(db *gorm.DB) *gorm.DB) (*[]models.Customer, error)
	GetDetail(id int) (*models.Customer, error)
	Insert(customer *models.Customer) error
	Update(req interface{}, id int) error
	Delete(id int) error
	Search(keyword string) (*[]models.Customer, error)
}

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		DB: db,
	}
}

func (r *customerRepository) FindAll(paginate func(db *gorm.DB) *gorm.DB) (*[]models.Customer, error) {
	var customers []models.Customer

	if err := r.DB.Scopes(paginate).Find(&customers).Error; err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r *customerRepository) GetDetail(id int) (*models.Customer, error) {
	var customer models.Customer
	if err := r.DB.Table("customers").Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Insert(customer *models.Customer) error {
	return r.DB.Table("customers").Create(&customer).Error
}

func (r *customerRepository) Update(req interface{}, id int) error {
	var customer models.Customer
	if err := r.DB.Table("customers").Clauses(clause.OnConflict{DoNothing: true}).Model(customer).Where("id = ?", id).Updates(req); err.Error != nil {
		return err.Error
	}
	r.DB.Model(&customer).Where("id = ?", id).Update("UpdatedAt", time.Now())
	return nil
}

func (r *customerRepository) Delete(id int) error {
	var customer models.Customer
	if res := r.DB.Table("customers").Clauses(clause.Returning{}).Where("id = ?", id).Delete(&customer); res.RowsAffected < 1 {
		return fmt.Errorf("customer does not exist")
	}
	return nil
}

func (r *customerRepository) Search(keyword string) (*[]models.Customer, error) {
	var customer []models.Customer
	if err := r.DB.Table("customers").Where("name LIKE ?", "%"+keyword+"%").Or("email LIKE ?", "%"+keyword+"%").Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
