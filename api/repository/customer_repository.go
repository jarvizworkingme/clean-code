package repository

import (
	"clean-code/api/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(db *gorm.DB, customer entity.Customer) (err error)
	FindByEmail(db *gorm.DB, email string) (entity.Customer, error)
}

type customerRepository struct {
	BaseRepository[entity.Customer]
}

func NewCustomerRepository() *customerRepository {
	return new(customerRepository)
}

func (u *customerRepository) Create(db *gorm.DB, customer entity.Customer) (err error) {
	return u.BaseRepository.Create(db, &customer)
}

func (u *customerRepository) FindByEmail(db *gorm.DB, email string) (customer entity.Customer, err error) {
	err = db.Where("email = ?", email).First(&customer).Error
	return
}
