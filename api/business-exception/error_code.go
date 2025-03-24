package business_exception

import "net/http"

const (
	ErrorGeneric           = "E400"
	ErrorDuplicateCustomer = "E451"
	ErrorInvalidAmount     = "E471"
	ErrorCustomerNotFound  = "E491"
	ErrorBillNotFound      = "E492"
)

var errorMessages = map[string]string{
	ErrorDuplicateCustomer: "Email already exists",
	ErrorCustomerNotFound:  "Customer not found",
	ErrorBillNotFound:      "Bill not found",
}

// by default, the status code is 422
var errorHttpStatus = map[string]int{
	ErrorDuplicateCustomer: http.StatusConflict,
	ErrorInvalidAmount:     http.StatusUnprocessableEntity,
}
