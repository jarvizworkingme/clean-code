package models

type GenericResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewSuccessResponse(code, message string) *GenericResponse {
	return &GenericResponse{
		Code:    code,
		Message: message,
	}
}
