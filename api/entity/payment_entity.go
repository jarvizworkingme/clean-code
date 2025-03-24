package entity

import "time"

func (u *Payment) TableName() string {
	return "payment"
}

type Payment struct {
	ID          int       `gorm:"primaryKey"`
	LoanID      int       `gorm:"type:integer;not null;index"`
	Loan        Loan      `gorm:"foreignKey:LoanID"`
	PaymentDate time.Time `gorm:"type:date;not null"`
	AmountPaid  float64   `gorm:"type:decimal(15,2);not null"`
	BaseEntity
}
