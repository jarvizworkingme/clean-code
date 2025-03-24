package repository

import (
	"clean-code/api/entity"
	"clean-code/api/models"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(db *gorm.DB, schedule entity.Schedule) error
	Update(db *gorm.DB, schedule entity.Schedule) error
	ListAll(db *gorm.DB, req models.ScheduleParams) (data []entity.Schedule, err error)
}

type scheduleRepository struct {
	BaseRepository[entity.Schedule]
}

func NewScheduleRepository() *scheduleRepository {
	return new(scheduleRepository)
}

func (u *scheduleRepository) Create(db *gorm.DB, schedule entity.Schedule) (err error) {
	return u.BaseRepository.Create(db, &schedule)
}

func (u *scheduleRepository) Update(db *gorm.DB, schedule entity.Schedule) error {
	return u.BaseRepository.Update(db, &schedule)
}

func (u *scheduleRepository) ListAll(db *gorm.DB, req models.ScheduleParams) (data []entity.Schedule, err error) {
	tx := db.Joins("Loan").Joins("Loan.Customer")
	if req.Email != "" {
		tx = tx.Where("email = ?", req.Email)
	}

	if req.Date != nil {
		tx = tx.Where("due_date <= ? and is_paid = ?", &req.Date, false)
	}

	if req.LoanID > 0 {
		tx = tx.Where("loan_id = ?", req.LoanID)
	}

	err = tx.Find(&data).Error
	return

}
