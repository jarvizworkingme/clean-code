package repository

import (
	"clean-code/api/entity"
	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateWithId(db *gorm.DB, loan entity.Loan) (id int, err error)
	Update(db *gorm.DB, loan entity.Loan) error
	ListAll(db *gorm.DB, email string) (data []entity.Loan, err error)
}

type loanRepository struct {
	BaseRepository[entity.Loan]
}

func NewLoanRepository() *loanRepository {
	return new(loanRepository)
}

func (u *loanRepository) CreateWithId(db *gorm.DB, loan entity.Loan) (id int, err error) {
	return u.BaseRepository.CreateWithID(db, &loan)
}

func (u *loanRepository) Update(db *gorm.DB, loan entity.Loan) error {
	return u.BaseRepository.Update(db, &loan)
}

func (u *loanRepository) ListAll(db *gorm.DB, email string) (data []entity.Loan, err error) {
	tx := db.Preload("Customer")
	if email != "" {
		tx = tx.Where("customer.email = ?", email)
	}
	err = tx.Find(&data).Order("id desc").Error
	return
}
