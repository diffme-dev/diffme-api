package errors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

type ApiError struct {
	Message    string   `json:"message"`
	Data       struct{} `json:"data"`
	StatusCode int      `json:"status_code"`
}

func NewApiError(c *fiber.Ctx, err error, statusCode int, data struct{}) error {
	apiError := ApiError{
		Message:    err.Error(),
		StatusCode: statusCode,
		Data:       data,
	}

	return c.Status(apiError.StatusCode).JSON(apiError)
}
