package converters

import (
	"clean-code/api/entity"
	"clean-code/api/models"
)

func ToSchedulesPaymentResponse(entity []entity.Schedule) []models.ScheduleResponse {

	mapPayments := make(map[int][]models.SchedulePayment)
	mapLoan := make(map[int]models.ScheduleResponse)
	mapTotalDue := make(map[int]float64)

	for _, v := range entity {
		mapPayments[v.LoanID] = append(mapPayments[v.LoanID], models.SchedulePayment{
			ID:         v.ID,
			WeekNumber: v.WeekNumber,
			DueDate:    v.DueDate,
			AmountDue:  v.AmountDue,
		})
		mapTotalDue[v.LoanID] += v.AmountDue

		if _, ok := mapLoan[v.LoanID]; !ok {
			mapLoan[v.LoanID] = models.ScheduleResponse{
				LoanId:             v.LoanID,
				Email:              v.Loan.Customer.Email,
				Amount:             v.Loan.PrincipalAmount,
				Interest:           v.Loan.InterestAmount,
				TotalAmount:        v.Loan.TotalAmount,
				OutstandingBalance: v.Loan.OutstandingBalance,
			}
		}
	}

	var data = make([]models.ScheduleResponse, len(mapLoan))
	var index int
	for _, v := range mapLoan {
		v.SchedulePayment = mapPayments[v.LoanId]
		v.CountDelinquent = len(v.SchedulePayment)
		v.IsDelinquent = v.CountDelinquent > 1
		v.TotalAmountDue = mapTotalDue[v.LoanId]
		data[index] = v
		index++
	}
	return data
}
