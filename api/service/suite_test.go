package service_test

import (
	"clean-code/api/entity"
	"clean-code/api/models"
	"clean-code/api/repository"
	"clean-code/api/service"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	CustomerService service.CustomerService
	LoanService     service.LoanService
	ScheduleService service.ScheduleService
	PaymentService  service.PaymentService
}

func (suite *Suite) SetupTest() {
	mockDb, mockQ, _ := sqlmock.New()
	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	}), &gorm.Config{})
	mockQ.ExpectBegin()
	mockQ.ExpectCommit()

	customerRepositoryMock := new(repository.CustomerRepositoryMock)
	loanRepositoryMock := new(repository.LoanRepositoryMock)
	scheduleRepositoryMock := new(repository.ScheduleRepositoryMock)
	paymentRepositoryMock := new(repository.PaymentRepositoryMock)

	customerRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(nil)
	customerRepositoryMock.On("FindByEmail", mock.Anything, mock.Anything).Return(entity.Customer{
		ID:    1,
		Email: "dummy@dummy.dummy",
	}, nil)

	loanRepositoryMock.On("CreateWithId", mock.Anything, mock.Anything).Return(1, nil)
	loanRepositoryMock.On("Update", mock.Anything, mock.Anything).Return(nil)
	loanRepositoryMock.On("ListAll", mock.Anything, mock.Anything).Return([]entity.Loan{
		{
			ID:              1,
			CustomerID:      1,
			PrincipalAmount: 1000000,
			InterestRate:    10,
			DurationWeek:    1,
			InterestAmount:  100000,
			TotalAmount:     1100000,
		},
		{
			ID:              2,
			CustomerID:      1,
			PrincipalAmount: 1000000,
			InterestRate:    10,
			DurationWeek:    1,
			InterestAmount:  100000,
			TotalAmount:     1100000,
		},
	}, nil)
	scheduleRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(nil)

	schedules := make([]entity.Schedule, 10)
	for i := range schedules {
		schedules[i] = entity.Schedule{
			ID:     i + 1,
			LoanID: 1,
			Loan: entity.Loan{
				ID:              1,
				CustomerID:      1,
				PrincipalAmount: 1000000,
				InterestRate:    10,
				DurationWeek:    1,
				InterestAmount:  100000,
				TotalAmount:     1100000,
				Customer: entity.Customer{
					ID:    1,
					Email: "test@gmail.com",
				},
				BaseEntity: entity.BaseEntity{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			WeekNumber: 1,
			DueDate:    time.Now(),
			AmountDue:  1100000,
			IsPaid:     false,
			BaseEntity: entity.BaseEntity{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
	}

	scheduleRepositoryMock.On("ListAll", mock.Anything, mock.Anything).Return(schedules, nil)
	scheduleRepositoryMock.On("Update", mock.Anything, mock.Anything).Return(nil)
	paymentRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(nil)

	suite.CustomerService = service.NewCustomerService(db, customerRepositoryMock)
	suite.LoanService = service.NewLoanService(db, loanRepositoryMock, customerRepositoryMock, scheduleRepositoryMock)
	suite.ScheduleService = service.NewScheduleService(db, scheduleRepositoryMock)
	suite.PaymentService = service.NewPaymentService(db, scheduleRepositoryMock, paymentRepositoryMock, loanRepositoryMock)
}

func (suite *Suite) TestCustomerServiceCreate() {

	err := suite.CustomerService.Create(models.CustomerRequest{
		Email: "test@gmail.com",
	})

	if err != nil {
		suite.T().Errorf("Code: %v", err)
	}
	suite.T().Log("Test Customer Service Create Done")
}

func (suite *Suite) TestLoanServiceCreate() {
	response, err := suite.LoanService.Create(models.LoanRequest{
		Email:        "test@gmail.com",
		Amount:       1000000,
		InterestRate: 10,
		Tenor:        1,
	})
	if err != nil {
		suite.T().Errorf("Code: %v", err)
	}
	if response.TotalAmount != 1100000 {
		suite.T().Errorf("Total Amount is not match")
	}

	suite.T().Log("Test Loan Service Create Done")
}

func (suite *Suite) TestLoanServiceListOutstandingBalance() {
	response, err := suite.LoanService.ListOutstandingBalance("")
	if err != nil {
		suite.T().Errorf("Code: %v", err)
	}

	if response[0].TotalAmount != 1100000 {
		suite.T().Errorf("Total Amount is not match")
	}

	suite.T().Log("Test Loan Service List Outstanding Balance Done")
}

func (suite *Suite) TestScheduleServiceListAll() {
	_, err := suite.ScheduleService.ListSchedulePayment("", false)
	if err != nil {
		suite.T().Errorf("Code: %v", err)
	}

	suite.T().Log("Test Schedule Service List All Done")
}

func (suite *Suite) TestPaymentServiceCreate() {
	err := suite.PaymentService.Pay(models.PaymentRequest{
		LoanID: 1,
		Amount: 1100000 * 10,
	})
	if err != nil {
		suite.T().Errorf("Code: %v", err)
	}

	suite.T().Log("Test Payment Service Create Done")
}

func TestRunAllService(t *testing.T) {
	suite.Run(t, new(Suite))
}
