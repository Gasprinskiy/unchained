package global

import "errors"

var (
	// ErrNoData данные не найдены"
	ErrNoData = errors.New("NoData")
	// ErrInternalError внутряя ошибка
	ErrInternalError = errors.New("InternalError")
	// ErrPermissionDenied отказано в доступе
	ErrPermissionDenied = errors.New("PermissionDenied")
	// ErrInvalidParam не верные параметры
	ErrInvalidParam = errors.New("InvalidParam")
	// ErrExpired время вышло
	ErrExpired = errors.New("Expired")
)
