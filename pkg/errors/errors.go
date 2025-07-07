package errors

import "errors"

type AppError struct {
	Code         int    // Код ошибки (например, 404, 500 и т.д.)
	Message      string // Сообщение для пользователя
	Err          error  // Подробная информация об ошибке (можно не показывать пользователю)
	IsUserFacing bool   // Может ли ошибка быть показана пользователю
}

const (
	InternalServerError = "internal server error"
	BadRequest          = "bad request"
	NotFound            = "not_found"

	IncorrectClientDataCode = 400
	InternalServerErrorCode = 500
)

func NewAppError(httpCode int, message string, err error, isUserFacing bool) *AppError {
	return &AppError{
		Code:         httpCode,
		Message:      message,
		Err:          err,
		IsUserFacing: isUserFacing,
	}
}

var (
	ErrEmptyAction  = errors.New("action did not affect the data")
	ErrDataNotFound = errors.New("data not found")
	ErrEmptyData    = errors.New("empty data")
)
