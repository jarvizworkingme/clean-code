package handler

import (
	"clean-code/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ListHandler(scheduleService service.ScheduleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				_ = c.Error(err)
			}
		}()

		email := strings.ToLower(c.Query("email"))
		isDelinquent := strings.ToLower(c.Query("is_delinquent")) == "true"

		response, err := scheduleService.ListSchedulePayment(email, isDelinquent)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, response)

	}
}
