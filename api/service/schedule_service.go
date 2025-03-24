package service

import (
	"clean-code/api/models"
	"clean-code/api/models/converters"
	"clean-code/api/repository"
	"gorm.io/gorm"
	"time"
)

type ScheduleService interface {
	ListSchedulePayment(email string, isDelinquent bool) (response []models.ScheduleResponse, err error)
}

type scheduleService struct {
	Db                 *gorm.DB
	ScheduleRepository repository.ScheduleRepository
}

func NewScheduleService(db *gorm.DB, scheduleRepository repository.ScheduleRepository) *scheduleService {
	return &scheduleService{Db: db, ScheduleRepository: scheduleRepository}
}

func (u *scheduleService) ListSchedulePayment(email string, isDelinquent bool) (response []models.ScheduleResponse, err error) {

	params := models.ScheduleParams{
		Email: email,
	}

	dueDatePayment := time.Now()
	//dueDatePayment := time.Now().AddDate(0, 0, 14)

	if isDelinquent {
		dueDatePayment = dueDatePayment.AddDate(0, 0, -14)
	}
	params.Date = &dueDatePayment

	data, err := u.ScheduleRepository.ListAll(u.Db, params)
	if err != nil {
		return
	}

	response = converters.ToSchedulesPaymentResponse(data)
	return
}
