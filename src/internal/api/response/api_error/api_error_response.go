package api_error

import "github.com/gofiber/fiber/v2"

type ApiErrorResponse struct {
	Code    int    `json:"code"`
	Status  uint   `json:"status"`
	Message string `json:"message"`
}

func PageNotFound() *ApiErrorResponse {
	return &ApiErrorResponse{
		Code:    1,
		Status:  fiber.StatusNotFound,
		Message: "The page you are looking for does not exist",
	}
}
func BadRequest() *ApiErrorResponse {
	return &ApiErrorResponse{
		Code:    2,
		Status:  fiber.StatusBadRequest,
		Message: "Bad Request",
	}
}

func InternalServerError() *ApiErrorResponse {
	return &ApiErrorResponse{
		Code:    3,
		Status:  fiber.StatusInternalServerError,
		Message: "Internal Server Error",
	}
}
