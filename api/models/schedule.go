package models

import "time"

type Schedule struct {
}

type ScheduleResponse struct {
	LoanId             int               `json:"loan_id"`
	Email              string            `json:"email"`
	Amount             float64           `json:"amount"`
	Interest           float64           `json:"interest"`
	TotalAmount        float64           `json:"total_amount"`
	OutstandingBalance float64           `json:"outstanding_balance"`
	IsDelinquent       bool              `json:"is_delinquent"`
	CountDelinquent    int               `json:"count_delinquent"`
	TotalAmountDue     float64           `json:"total_amount_due"`
	SchedulePayment    []SchedulePayment `json:"schedule_payment"`
}

type SchedulePayment struct {
	ID         int       `json:"id"`
	WeekNumber int       `json:"week_number"`
	DueDate    time.Time `json:"due_date"`
	AmountDue  float64   `json:"amount_due"`
}

type ScheduleParams struct {
	Email  string
	Date   *time.Time
	LoanID int
}
