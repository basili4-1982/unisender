package errors

import (
	"fmt"
	"strings"
)

// APIError представляет ошибку API Unisender
type APIError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Result     string `json:"result"`
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("API error [%s]: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("API error: %s (status: %d)", e.Message, e.StatusCode)
}

// RateLimitError ошибка превышения лимита запросов
type RateLimitError struct {
	*APIError
	RetryAfter int // seconds
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit exceeded: %s (retry after %d seconds)", e.Message, e.RetryAfter)
}

// ValidationError ошибка валидации входных данных
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// ValidationErrors множественная ошибка валидации
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (e *ValidationErrors) Error() string {
	if len(e.Errors) == 0 {
		return "validation errors"
	}

	var errMsgs []string
	for _, err := range e.Errors {
		errMsgs = append(errMsgs, err.Error())
	}
	return fmt.Sprintf("validation errors: %s", strings.Join(errMsgs, "; "))
}

// Add добавляет ошибку валидации
func (e *ValidationErrors) Add(field, message string) {
	e.Errors = append(e.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// HasErrors возвращает true, если есть ошибки
func (e *ValidationErrors) HasErrors() bool {
	return len(e.Errors) > 0
}

// AuthenticationError ошибка аутентификации
type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("authentication failed: %s", e.Message)
}

// NotFoundError ошибка "не найдено"
type NotFoundError struct {
	Resource string
	ID       string
	Message  string
}

func (e *NotFoundError) Error() string {
	if e.Resource != "" && e.ID != "" {
		return fmt.Sprintf("%s with ID '%s' not found", e.Resource, e.ID)
	}
	if e.Message != "" {
		return e.Message
	}
	return "resource not found"
}

// InternalError внутренняя ошибка сервера
type InternalError struct {
	Message string
	Err     error
}

func (e *InternalError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("internal error: %s: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("internal error: %s", e.Message)
}

func (e *InternalError) Unwrap() error {
	return e.Err
}

// IsRetryable проверяет, можно ли повторить запрос
func IsRetryable(err error) bool {
	switch e := err.(type) {
	case *APIError:
		// 429 - Too Many Requests, 5xx - Server Errors
		return e.StatusCode == 429 || e.StatusCode >= 500
	case *RateLimitError:
		return true
	case *InternalError:
		return true
	default:
		return false
	}
}

// IsValidationError проверяет, является ли ошибка ошибкой валидации
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	if ok {
		return true
	}
	_, ok = err.(*ValidationErrors)
	return ok
}

// IsNotFoundError проверяет, является ли ошибка ошибкой "не найдено"
func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// IsRateLimitError проверяет, является ли ошибка ошибкой rate limit
func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

// IsAuthenticationError проверяет, является ли ошибка ошибкой аутентификации
func IsAuthenticationError(err error) bool {
	_, ok := err.(*AuthenticationError)
	return ok
}
