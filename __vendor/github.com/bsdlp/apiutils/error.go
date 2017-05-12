package apiutils

import "net/http"

// Error is an error
type Error interface {
	StatusCode() int
	Error() string
}

type err struct {
	status int
	msg    string
}

func (e err) Error() string { return e.msg }

// StatusCode returns the recommended http status code for this error
func (e err) StatusCode() int { return e.status }

// NewError returns an Error
func NewError(statusCode int, msg string) Error {
	if msg == "" {
		msg = http.StatusText(statusCode)
	}
	return err{
		status: statusCode,
		msg:    msg,
	}
}

// default http errors
var (
	ErrBadRequest                    = NewError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	ErrUnauthorized                  = NewError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	ErrPaymentRequired               = NewError(http.StatusPaymentRequired, http.StatusText(http.StatusPaymentRequired))
	ErrForbidden                     = NewError(http.StatusForbidden, http.StatusText(http.StatusForbidden))
	ErrNotFound                      = NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	ErrMethodNotAllowed              = NewError(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	ErrNotAcceptable                 = NewError(http.StatusNotAcceptable, http.StatusText(http.StatusNotAcceptable))
	ErrProxyAuthRequired             = NewError(http.StatusProxyAuthRequired, http.StatusText(http.StatusProxyAuthRequired))
	ErrRequestTimeout                = NewError(http.StatusRequestTimeout, http.StatusText(http.StatusRequestTimeout))
	ErrConflict                      = NewError(http.StatusConflict, http.StatusText(http.StatusConflict))
	ErrGone                          = NewError(http.StatusGone, http.StatusText(http.StatusGone))
	ErrLengthRequired                = NewError(http.StatusLengthRequired, http.StatusText(http.StatusLengthRequired))
	ErrPreconditionFailed            = NewError(http.StatusPreconditionFailed, http.StatusText(http.StatusPreconditionFailed))
	ErrRequestEntityTooLarge         = NewError(http.StatusRequestEntityTooLarge, http.StatusText(http.StatusRequestEntityTooLarge))
	ErrRequestURITooLong             = NewError(http.StatusRequestURITooLong, http.StatusText(http.StatusRequestURITooLong))
	ErrUnsupportedMediaType          = NewError(http.StatusUnsupportedMediaType, http.StatusText(http.StatusUnsupportedMediaType))
	ErrRequestedRangeNotSatisfiable  = NewError(http.StatusRequestedRangeNotSatisfiable, http.StatusText(http.StatusRequestedRangeNotSatisfiable))
	ErrExpectationFailed             = NewError(http.StatusExpectationFailed, http.StatusText(http.StatusExpectationFailed))
	ErrTeapot                        = NewError(http.StatusTeapot, http.StatusText(http.StatusTeapot))
	ErrUnprocessableEntity           = NewError(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))
	ErrLocked                        = NewError(http.StatusLocked, http.StatusText(http.StatusLocked))
	ErrFailedDependency              = NewError(http.StatusFailedDependency, http.StatusText(http.StatusFailedDependency))
	ErrUpgradeRequired               = NewError(http.StatusUpgradeRequired, http.StatusText(http.StatusUpgradeRequired))
	ErrPreconditionRequired          = NewError(http.StatusPreconditionRequired, http.StatusText(http.StatusPreconditionRequired))
	ErrTooManyRequests               = NewError(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
	ErrRequestHeaderFieldsTooLarge   = NewError(http.StatusRequestHeaderFieldsTooLarge, http.StatusText(http.StatusRequestHeaderFieldsTooLarge))
	ErrUnavailableForLegalReasons    = NewError(http.StatusUnavailableForLegalReasons, http.StatusText(http.StatusUnavailableForLegalReasons))
	ErrInternalServerError           = NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	ErrNotImplemented                = NewError(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
	ErrBadGateway                    = NewError(http.StatusBadGateway, http.StatusText(http.StatusBadGateway))
	ErrServiceUnavailable            = NewError(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable))
	ErrGatewayTimeout                = NewError(http.StatusGatewayTimeout, http.StatusText(http.StatusGatewayTimeout))
	ErrHTTPVersionNotSupported       = NewError(http.StatusHTTPVersionNotSupported, http.StatusText(http.StatusHTTPVersionNotSupported))
	ErrVariantAlsoNegotiates         = NewError(http.StatusVariantAlsoNegotiates, http.StatusText(http.StatusVariantAlsoNegotiates))
	ErrInsufficientStorage           = NewError(http.StatusInsufficientStorage, http.StatusText(http.StatusInsufficientStorage))
	ErrLoopDetected                  = NewError(http.StatusLoopDetected, http.StatusText(http.StatusLoopDetected))
	ErrNotExtended                   = NewError(http.StatusNotExtended, http.StatusText(http.StatusNotExtended))
	ErrNetworkAuthenticationRequired = NewError(http.StatusNetworkAuthenticationRequired, http.StatusText(http.StatusNetworkAuthenticationRequired))
)
