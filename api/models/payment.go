package models

type PaymentRequest struct {
	Amount float64 `json:"amount"`
	LoanID int     `json:"loan_id"`
}
