package errors

import (
	"errors"
)

type AppError struct {
	Code         int    // Код ошибки (например, 404, 500 и т.д.)
	Message      string // Сообщение для пользователя
	Err          error  // Подробная информация об ошибке (можно не показывать пользователю)
	IsUserFacing bool   // Может ли ошибка быть показана пользователю
}

type DBError struct {
	Message string
	Err     error
}

const (
	InternalServerError = "internal server error"
	BadRequest          = "bad request"
	NotFound            = "not_found"

	IncorrectClientDataCode = 400
	InternalServerErrorCode = 500
	NotFoundErrorCode       = 404
)

func NewAppError(httpCode int, message string, err error, isUserFacing bool) *AppError {
	return &AppError{
		Code:         httpCode,
		Message:      message,
		Err:          err,
		IsUserFacing: isUserFacing,
	}
}

func NewDBError(message string, dbError error) *AppError {

	return &AppError{
		Code:         InternalServerErrorCode,
		Message:      message,
		Err:          dbError,
		IsUserFacing: false,
	}
}

var (
	ErrEmptyAction  = errors.New("action did not affect the data")
	ErrDataNotFound = errors.New("data not found")
	ErrEmptyData    = errors.New("empty data")
)

func Is(err any, err2 error) bool {
	if e, ok := err.(error); ok {
		return errors.Is(e, err2)
	}
	return false
}
