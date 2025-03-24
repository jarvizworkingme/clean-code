package service

type Service struct {
	CustomerService CustomerService
	LoanService     LoanService
	ScheduleService ScheduleService
	PaymentService  PaymentService
}
