package entity

import "time"

func (u *Loan) TableName() string {
	return "loan"
}

type Loan struct {
	ID                 int       `gorm:"primaryKey"`
	CustomerID         int       `gorm:"type:integer;not null;index"`
	Customer           Customer  `gorm:"foreignKey:CustomerID"`
	PrincipalAmount    float64   `gorm:"type:decimal(15,2);not null"`
	InterestRate       float64   `gorm:"type:decimal(5,2);not null"`
	InterestAmount     float64   `gorm:"type:decimal(15,2);not null"`
	TotalAmount        float64   `gorm:"type:decimal(15,2);not null"`
	OutstandingBalance float64   `gorm:"type:decimal(15,2);not null"`
	DurationWeek       int       `gorm:"type:integer;not null"`
	StartDate          time.Time `gorm:"type:date;not null"`
	BaseEntity
}
