package entity

import "time"

func (u *Schedule) TableName() string {
	return "schedule"
}

type Schedule struct {
	ID         int       `gorm:"primaryKey"`
	LoanID     int       `gorm:"type:integer;not null;index"`
	Loan       Loan      `gorm:"foreignKey:LoanID"`
	WeekNumber int       `gorm:"type:integer;not null"`
	DueDate    time.Time `gorm:"type:date;not null"`
	AmountDue  float64   `gorm:"type:decimal(15,2);not null"`
	IsPaid     bool      `gorm:"type:boolean;not null"`
	BaseEntity
}
