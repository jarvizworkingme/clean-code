package models

import "time"

type LoanResponse struct {
	ID                 int       `json:"id"`
	Amount             float64   `json:"amount"`
	Email              string    `json:"email"`
	Interest           float64   `json:"interest"`
	TotalAmount        float64   `json:"total_amount"`
	OutstandingBalance float64   `gorm:"type:decimal(15,2);not null"`
	StartDate          time.Time `json:"start_date"`
}

type LoanRequest struct {
	Email        string  `json:"email"`
	InterestRate float64 `json:"interest_rate"`
	Amount       float64 `json:"amount"`
	Tenor        int     `json:"tenor"`
}
