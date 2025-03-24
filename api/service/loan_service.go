package service

import (
	business_exception "clean-code/api/business-exception"
	"clean-code/api/entity"
	"clean-code/api/models"
	"clean-code/api/models/converters"
	"clean-code/api/repository"
	"errors"
	"gorm.io/gorm"
	"time"
)

type LoanService interface {
	Create(req models.LoanRequest) (response models.LoanResponse, err error)
	ListOutstandingBalance(email string) (response []models.LoanResponse, err error)
}

type loanService struct {
	Db                 *gorm.DB
	LoanRepository     repository.LoanRepository
	CustomerRepository repository.CustomerRepository
	ScheduleRepository repository.ScheduleRepository
}

func NewLoanService(db *gorm.DB, loanRepository repository.LoanRepository, customerRepository repository.CustomerRepository, scheduleRepository repository.ScheduleRepository) *loanService {
	return &loanService{Db: db, LoanRepository: loanRepository, CustomerRepository: customerRepository, ScheduleRepository: scheduleRepository}
}

func (u *loanService) Create(req models.LoanRequest) (response models.LoanResponse, err error) {
	err = u.Db.Transaction(func(tx *gorm.DB) error {
		customer, err := u.CustomerRepository.FindByEmail(tx, req.Email)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return business_exception.NewBusinessError(business_exception.ErrorCustomerNotFound, "")
			}
			return err
		}

		var loan entity.Loan
		loan.CustomerID = customer.ID
		loan.Customer = customer
		loan.PrincipalAmount = req.Amount
		loan.InterestRate = req.InterestRate
		loan.DurationWeek = req.Tenor

		loan.InterestAmount = loan.PrincipalAmount * loan.InterestRate / 100
		loan.TotalAmount = loan.PrincipalAmount + loan.InterestAmount
		loan.OutstandingBalance = loan.TotalAmount
		loan.StartDate = time.Now()

		loan.ID, err = u.LoanRepository.CreateWithId(tx, loan)
		if err != nil {
			return err
		}

		var schedule entity.Schedule
		schedule.LoanID = loan.ID
		schedule.IsPaid = false
		schedule.AmountDue = loan.TotalAmount / float64(loan.DurationWeek)

		for i := 1; i <= req.Tenor; i++ {
			schedule.DueDate = loan.StartDate.AddDate(0, 0, i*7)
			schedule.WeekNumber = i
			err = u.ScheduleRepository.Create(tx, schedule)
			if err != nil {
				return err
			}
		}
		response = converters.ToLoanResponse(loan)

		return err
	})
	return response, err
}

func (u *loanService) ListOutstandingBalance(email string) (response []models.LoanResponse, err error) {

	loans, err := u.LoanRepository.ListAll(u.Db, email)
	if err != nil {
		return response, err
	}
	return converters.ToLoansResponse(loans), err
}
