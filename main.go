package main

import (
	"clean-code/api"
	"clean-code/api/entity"
	"clean-code/api/repository"
	"clean-code/api/service"
	"clean-code/api/storages"
	"clean-code/config"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {

	//Postgresql
	db := storages.NewDbConnection(config.NewConfigPostgresql())
	if strings.ToLower(os.Getenv("MIGRATE_DB")) == "true" {
		logrus.Println("Drop and create DB")
		entity.DropAndCreateDb(db)
	}

	//repository
	customerRepository := repository.NewCustomerRepository()
	loanRepository := repository.NewLoanRepository()
	scheduleRepository := repository.NewScheduleRepository()
	paymentRepository := repository.NewPaymentRepository()

	s := service.Service{
		CustomerService: service.NewCustomerService(db, customerRepository),
		LoanService:     service.NewLoanService(db, loanRepository, customerRepository, scheduleRepository),
		ScheduleService: service.NewScheduleService(db, scheduleRepository),
		PaymentService:  service.NewPaymentService(db, scheduleRepository, paymentRepository, loanRepository),
	}

	if err := api.NewRoutes(s).Run(":7778"); err != nil {
		logrus.Errorln(err)
	}
}
