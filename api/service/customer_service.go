package service

import (
	business_exception "clean-code/api/business-exception"
	"clean-code/api/entity"
	"clean-code/api/models"
	"clean-code/api/repository"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"strings"
)

type CustomerService interface {
	Create(req models.CustomerRequest) (err error)
}

type customerService struct {
	Db                 *gorm.DB
	CustomerRepository repository.CustomerRepository
}

func NewCustomerService(db *gorm.DB, customerRepository repository.CustomerRepository) *customerService {
	return &customerService{Db: db, CustomerRepository: customerRepository}
}

func (u *customerService) Create(req models.CustomerRequest) (err error) {
	err = u.CustomerRepository.Create(u.Db, entity.Customer{
		Email: strings.ToLower(req.Email),
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return business_exception.NewBusinessError(business_exception.ErrorDuplicateCustomer, "")
		}
	}
	return
}
