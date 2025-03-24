package repository

import (
	"clean-code/api/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreateWithId(db *gorm.DB, payment entity.Payment) (id int, err error)
	Create(db *gorm.DB, payment entity.Payment) error
	Update(db *gorm.DB, payment entity.Payment) error
}

type paymentRepository struct {
	BaseRepository[entity.Payment]
}

func NewPaymentRepository() *paymentRepository {
	return new(paymentRepository)
}

func (u *paymentRepository) Create(db *gorm.DB, payment entity.Payment) error {
	return u.BaseRepository.Create(db, &payment)
}

func (u *paymentRepository) CreateWithId(db *gorm.DB, payment entity.Payment) (id int, err error) {
	return u.BaseRepository.CreateWithID(db, &payment)
}

func (u *paymentRepository) Update(db *gorm.DB, payment entity.Payment) error {
	return u.BaseRepository.Update(db, &payment)
}
