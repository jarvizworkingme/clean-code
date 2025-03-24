package service

import (
	business_exception "clean-code/api/business-exception"
	"clean-code/api/entity"
	"clean-code/api/models"
	"clean-code/api/repository"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type PaymentService interface {
	Pay(req models.PaymentRequest) (err error)
}

type paymentService struct {
	Db                 *gorm.DB
	ScheduleRepository repository.ScheduleRepository
	PaymentRepository  repository.PaymentRepository
	LoanRepository     repository.LoanRepository
}

func NewPaymentService(db *gorm.DB, scheduleRepository repository.ScheduleRepository, paymentRepository repository.PaymentRepository, loanRepository repository.LoanRepository) *paymentService {
	return &paymentService{Db: db, ScheduleRepository: scheduleRepository, PaymentRepository: paymentRepository, LoanRepository: loanRepository}
}

func (u *paymentService) Pay(req models.PaymentRequest) (err error) {

	//now := time.Now().AddDate(0, 0, 14)
	now := time.Now()
	schedulePayment, err := u.ScheduleRepository.ListAll(u.Db, models.ScheduleParams{
		Date:   &now,
		LoanID: req.LoanID,
	})

	if err != nil {
		return
	}

	if len(schedulePayment) == 0 {
		return business_exception.NewBusinessError(business_exception.ErrorBillNotFound, "")
	}

	var totalAmount float64
	for _, v := range schedulePayment {
		totalAmount += v.AmountDue
	}

	if totalAmount != req.Amount {
		return business_exception.NewBusinessError(business_exception.ErrorInvalidAmount, fmt.Sprint("Invalid amount, expected : ", totalAmount))
	}

	err = u.Db.Transaction(func(tx *gorm.DB) (err error) {
		var loan entity.Loan
		for _, v := range schedulePayment {
			v.IsPaid = true
			err = u.ScheduleRepository.Update(tx, v)
			if err != nil {
				return
			}
			if loan.ID == 0 {
				loan = v.Loan
			}
		}

		err = u.PaymentRepository.Create(tx, entity.Payment{
			LoanID:      req.LoanID,
			PaymentDate: now,
			AmountPaid:  req.Amount,
		})

		if err != nil {
			return
		}

		loan.OutstandingBalance -= req.Amount
		err = u.LoanRepository.Update(tx, loan)

		return err
	})

	return
}
