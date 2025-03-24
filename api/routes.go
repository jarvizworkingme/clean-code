package api

import (
	"clean-code/api/handler"
	"clean-code/api/middleware"
	"clean-code/api/service"
	"github.com/gin-gonic/gin"
)

func NewRoutes(service service.Service) (r *gin.Engine) {
	r = gin.Default()
	r.Use(middleware.BusinessError())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)
		v1.POST("/customer", handler.CreateCustomerHandler(service.CustomerService))
		v1.POST("/loan", handler.CreateLoanHandler(service.LoanService))
		v1.POST("/payment", handler.PaymentHandler(service.PaymentService))
		v1.GET("/out-standing", handler.ListOutStanding(service.LoanService))
		v1.GET("/payment/schedule", handler.ListHandler(service.ScheduleService))
	}

	return r
}
