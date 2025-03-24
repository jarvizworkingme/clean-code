package handler

import (
	"clean-code/api/models"
	"clean-code/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateLoanHandler(loanService service.LoanService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				_ = c.Error(err)
			}
		}()

		var req models.LoanRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			return
		}

		response, err := loanService.Create(req)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func ListOutStanding(loanService service.LoanService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		response, err := loanService.ListOutstandingBalance(email)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
