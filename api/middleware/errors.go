package middleware

import (
	businessexception "clean-code/api/business-exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BusinessError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			lastError := c.Errors.Last()
			if lastError.Type == gin.ErrorTypePublic {
				if businessException, ok := lastError.Meta.(businessexception.Error); ok {
					c.JSON(businessException.HttpStatus, businessException)
					return
				}
			}

			c.JSON(http.StatusInternalServerError, businessexception.NewBusinessError(businessexception.ErrorGeneric, lastError.Error()))
		}
	}
}
