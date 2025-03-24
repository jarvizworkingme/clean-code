package handler

import (
	"clean-code/api/models"
	"clean-code/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PaymentHandler(paymentService service.PaymentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				_ = c.Error(err)
			}
		}()

		var req models.PaymentRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			return
		}

		err = paymentService.Pay(req)
		if err != nil {
			return
		}

		c.Status(http.StatusCreated)
	}
}
