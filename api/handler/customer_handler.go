package handler

import (
	"clean-code/api/models"
	"clean-code/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCustomerHandler(customerService service.CustomerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				_ = c.Error(err)
			}
		}()

		var req models.CustomerRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			return
		}

		err = customerService.Create(req)
		if err != nil {
			return
		}

		c.Status(http.StatusCreated)
	}
}
