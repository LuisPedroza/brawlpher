package api

import "net/http"

type Error struct {
	Message    string
	StatusCode int
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrBadRequest          = Error{Message: "bad request", StatusCode: http.StatusBadRequest}
	ErrForbidden           = Error{Message: "forbidden", StatusCode: http.StatusForbidden}
	ErrNotFound            = Error{Message: "not found", StatusCode: http.StatusNotFound}
	ErrThrottled           = Error{Message: "throttled", StatusCode: http.StatusTooManyRequests}
	ErrInternalServerError = Error{Message: "internal server error", StatusCode: http.StatusInternalServerError}
	ErrServerUnavailable   = Error{Message: "server unavailable", StatusCode: http.StatusServiceUnavailable}

	StatusToError = map[int]Error{
		http.StatusBadRequest:          ErrBadRequest,
		http.StatusForbidden:           ErrForbidden,
		http.StatusNotFound:            ErrNotFound,
		http.StatusTooManyRequests:     ErrThrottled,
		http.StatusInternalServerError: ErrInternalServerError,
		http.StatusServiceUnavailable:  ErrServerUnavailable,
	}
)
