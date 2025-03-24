package converters

import (
	"clean-code/api/entity"
	"clean-code/api/models"
)

func ToLoanResponse(entity entity.Loan) models.LoanResponse {
	return models.LoanResponse{
		ID:                 entity.ID,
		Email:              entity.Customer.Email,
		Amount:             entity.PrincipalAmount,
		Interest:           entity.InterestAmount,
		TotalAmount:        entity.TotalAmount,
		OutstandingBalance: entity.OutstandingBalance,
		StartDate:          entity.StartDate,
	}
}

func ToLoansResponse(entity []entity.Loan) []models.LoanResponse {
	var data = make([]models.LoanResponse, len(entity))
	for i, v := range entity {
		data[i] = ToLoanResponse(v)
	}
	return data
}
