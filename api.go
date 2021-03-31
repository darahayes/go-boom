package boom

import "net/http"

// BadRequest responds with a 400 Bad Request error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderBadRequest(w http.ResponseWriter, message ...interface{}) {
	Render(w, BadRequest(message...))
}
func BadRequest(message ...interface{}) Err {
	return Boomify(http.StatusBadRequest, message...)
}

// Unathorized responds with a 401 Unauthorized error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderUnathorized(w http.ResponseWriter, message ...interface{}) {
	Render(w, Unathorized(message...))
}
func Unathorized(message ...interface{}) Err {
	return Boomify(http.StatusUnauthorized, message...)
}

// PaymentRequired responds with a 402 Payment Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderPaymentRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, PaymentRequired(message...))
}
func PaymentRequired(message ...interface{}) Err {
	return Boomify(http.StatusPaymentRequired, message...)
}

// Forbidden responds with a 403 Forbidden error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderForbidden(w http.ResponseWriter, message ...interface{}) {
	Render(w, Forbidden(message...))
}
func Forbidden(message ...interface{}) Err {
	return Boomify(http.StatusForbidden, message...)
}

// NotFound responds with a 404 Not Found error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderNotFound(w http.ResponseWriter, message ...interface{}) {
	Render(w, NotFound(message...))
}
func NotFound(message ...interface{}) Err {
	return Boomify(http.StatusNotFound, message...)
}

// MethodNotAllowed responds with a 405 Method Not Allowed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderMethodNotAllowed(w http.ResponseWriter, message ...interface{}) {
	Render(w, MethodNotAllowed(message...))
}
func MethodNotAllowed(message ...interface{}) Err {
	return Boomify(http.StatusMethodNotAllowed, message...)
}

// NotAcceptable responds with a 406 Not Acceptable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderNotAcceptable(w http.ResponseWriter, message ...interface{}) {
	Render(w, NotAcceptable(message...))
}
func NotAcceptable(message ...interface{}) Err {
	return Boomify(http.StatusNotAcceptable, message...)
}

// ProxyAuthRequired responds with a 407 Proxy Authentication Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderProxyAuthRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, ProxyAuthRequired(message...))
}
func ProxyAuthRequired(message ...interface{}) Err {
	return Boomify(http.StatusProxyAuthRequired, message...)
}

// RequestTimeout responds with a 408 Request Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderRequestTimeout(w http.ResponseWriter, message ...interface{}) {
	Render(w, RequestTimeout(message...))
}
func RequestTimeout(message ...interface{}) Err {
	return Boomify(http.StatusRequestTimeout, message...)
}

// Conflict responds with a 409 Conflict error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderConflict(w http.ResponseWriter, message ...interface{}) {
	Render(w, Conflict(message...))
}
func Conflict(message ...interface{}) Err {
	return Boomify(http.StatusConflict, message...)
}

// Gone responds with a 410 Gone error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderGone(w http.ResponseWriter, message ...interface{}) {
	Render(w, Gone(message...))
}
func Gone(message ...interface{}) Err {
	return Boomify(http.StatusGone, message...)
}

// LengthRequired responds with a 411 Length Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderLengthRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, LengthRequired(message...))
}
func LengthRequired(message ...interface{}) Err {
	return Boomify(http.StatusLengthRequired, message...)
}

// PreconditionFailed responds with a 412 Precondition Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderPreconditionFailed(w http.ResponseWriter, message ...interface{}) {
	Render(w, PreconditionFailed(message...))
}
func PreconditionFailed(message ...interface{}) Err {
	return Boomify(http.StatusPreconditionFailed, message...)
}

// EntityTooLarge responds with a 413 Request Entity Too Large error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderEntityTooLarge(w http.ResponseWriter, message ...interface{}) {
	Render(w, EntityTooLarge(message...))
}
func EntityTooLarge(message ...interface{}) Err {
	return Boomify(http.StatusRequestEntityTooLarge, message...)
}

// URITooLong responds with a 414 Request-URI Too Large error
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderURITooLong(w http.ResponseWriter, message ...interface{}) {
	Render(w, URITooLong(message...))
}
func URITooLong(message ...interface{}) Err {
	return Boomify(http.StatusRequestURITooLong, message...)
}

// UnsupportedMediaType responds with a 415 Unsupported Media Type error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderUnsupportedMediaType(w http.ResponseWriter, message ...interface{}) {
	Render(w, UnsupportedMediaType(message...))
}
func UnsupportedMediaType(message ...interface{}) Err {
	return Boomify(http.StatusUnsupportedMediaType, message...)
}

// RangeNotSatisfiable responds with a 416 Requested Range Not Satisfiable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderRangeNotSatisfiable(w http.ResponseWriter, message ...interface{}) {
	Render(w, RangeNotSatisfiable(message...))
}
func RangeNotSatisfiable(message ...interface{}) Err {
	return Boomify(http.StatusRequestedRangeNotSatisfiable, message...)
}

// ExpectationFailed responds with a 417 Expectation Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderExpectationFailed(w http.ResponseWriter, message ...interface{}) {
	Render(w, ExpectationFailed(message...))
}
func ExpectationFailed(message ...interface{}) Err {
	return Boomify(http.StatusExpectationFailed, message...)
}

// Teapot responds with a 418 I'm a Teapot error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderTeapot(w http.ResponseWriter, message ...interface{}) {
	Render(w, Teapot(message...))
}
func Teapot(message ...interface{}) Err {
	return Boomify(http.StatusTeapot, message...)
}

// UnprocessableEntity responds with a 422 Unprocessable Entity error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderUnprocessableEntity(w http.ResponseWriter, message ...interface{}) {
	Render(w, UnprocessableEntity(message...))
}
func UnprocessableEntity(message ...interface{}) Err {
	return Boomify(http.StatusUnprocessableEntity, message...)
}

// Locked responds with a 423 Locked error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderLocked(w http.ResponseWriter, message ...interface{}) {
	Render(w, Locked(message...))
}
func Locked(message ...interface{}) Err {
	return Boomify(http.StatusLocked, message...)
}

// PreconditionRequired responds with a 428 Precondition Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderPreconditionRequired(w http.ResponseWriter, message ...interface{}) {
	Render(w, PreconditionRequired(message...))
}
func PreconditionRequired(message ...interface{}) Err {
	return Boomify(http.StatusPreconditionRequired, message...)
}

// TooManyRequests responds with a 429 Too Many Requests error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderTooManyRequests(w http.ResponseWriter, message ...interface{}) {
	Render(w, TooManyRequests(message...))
}
func TooManyRequests(message ...interface{}) Err {
	return Boomify(http.StatusTooManyRequests, message...)
}

// TooManyRequests responds with a 431 request header fields too large error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderHeaderFieldsTooLarge(w http.ResponseWriter, message ...interface{}) {
	Render(w, HeaderFieldsTooLarge(message...))
}

func HeaderFieldsTooLarge(message ...interface{}) Err {
	return Boomify(http.StatusRequestHeaderFieldsTooLarge, message...)
}

// Illegal responds with a 451 Unavailable For Legal Reasons error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderIllegal(w http.ResponseWriter, message ...interface{}) {
	Render(w, Illegal(message...))
}
func Illegal(message ...interface{}) Err {
	return Boomify(http.StatusUnavailableForLegalReasons, message...)
}

// InternalServerError responds with a 500 Internal Server Error error. Alias for boom.Internal.
// Takes an optional message of either type string or type error
// but, will always return a generic message in the response body.
// This behaviour protects the developer from accidentally returning
// sensitive data in the response during a panic.
func RenderInternal(w http.ResponseWriter, message ...interface{}) {
	Render(w, Internal())
}
func Internal(message ...interface{}) Err {
	return Boomify(http.StatusInternalServerError)
}

func RenderBadImplementation(w http.ResponseWriter, message ...interface{}) {
	RenderInternal(w)
}

func BadImplementation(message ...interface{}) Err {
	return Internal(message...)
}

// NotImplemented responds with a 501 Not Implemented error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderNotImplemented(w http.ResponseWriter, message ...interface{}) {
	Render(w, NotImplemented(message...))
}
func NotImplemented(message ...interface{}) Err {
	return Boomify(http.StatusNotImplemented, message...)
}

// BadGateway responds with a 502 Bad Gateway error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderBadGateway(w http.ResponseWriter, message ...interface{}) {
	Render(w, BadGateway(message...))
}
func BadGateway(message ...interface{}) Err {
	return Boomify(http.StatusBadGateway, message...)
}

// ServiceUnavailable .eturns a 503 Service Unavailable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderServiceUnavailable(w http.ResponseWriter, message ...interface{}) {
	Render(w, ServiceUnavailable(message...))
}
func ServiceUnavailable(message ...interface{}) Err {
	return Boomify(http.StatusServiceUnavailable, message...)
}

// GatewayTimeout responds with a 504 Gateway Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RenderGatewayTimeout(w http.ResponseWriter, message ...interface{}) {
	Render(w, GatewayTimeout(message...))
}
func GatewayTimeout(message ...interface{}) Err {
	return Boomify(http.StatusGatewayTimeout, message...)
}
