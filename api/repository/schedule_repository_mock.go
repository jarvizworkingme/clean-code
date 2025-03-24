package repository

import (
	"clean-code/api/entity"
	"clean-code/api/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ScheduleRepositoryMock struct {
	mock.Mock
}

func (s *ScheduleRepositoryMock) Create(db *gorm.DB, schedule entity.Schedule) error {
	arguments := s.Called(schedule)
	return arguments.Error(0)
}

func (s *ScheduleRepositoryMock) Update(db *gorm.DB, schedule entity.Schedule) error {
	arguments := s.Called(schedule)
	return arguments.Error(0)
}

func (s *ScheduleRepositoryMock) ListAll(db *gorm.DB, req models.ScheduleParams) ([]entity.Schedule, error) {
	arguments := s.Called(req)
	return arguments.Get(0).([]entity.Schedule), arguments.Error(1)
}
