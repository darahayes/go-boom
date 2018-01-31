package boom

import "net/http"

// BadRequest responds with a 400 Bad Request error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func BadRequest(w http.ResponseWriter, message ...interface{}) {
	boom(w, 400, message...)
}

// Unathorized responds with a 401 Unauthorized error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func Unathorized(w http.ResponseWriter, message ...interface{}) {
	boom(w, 401, message...)
}

// PaymentRequired responds with a 402 Payment Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func PaymentRequired(w http.ResponseWriter, message ...interface{}) {
	boom(w, 402, message...)
}

// Forbidden responds with a 403 Forbidden error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func Forbidden(w http.ResponseWriter, message ...interface{}) {
	boom(w, 403, message...)
}

// NotFound responds with a 404 Not Found error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func NotFound(w http.ResponseWriter, message ...interface{}) {
	boom(w, 404, message...)
}

// MethodNotAllowed responds with a 405 Method Not Allowed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func MethodNotAllowed(w http.ResponseWriter, message ...interface{}) {
	boom(w, 405, message...)
}

// NotAcceptable responds with a 406 Not Acceptable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func NotAcceptable(w http.ResponseWriter, message ...interface{}) {
	boom(w, 406, message...)
}

// ProxyAuthRequired responds with a 407 Proxy Authentication Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func ProxyAuthRequired(w http.ResponseWriter, message ...interface{}) {
	boom(w, 407, message...)
}

// ClientTimeout responds with a 408 Request Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func ClientTimeout(w http.ResponseWriter, message ...interface{}) {
	boom(w, 408, message...)
}

// Conflict responds with a 409 Conflict error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func Conflict(w http.ResponseWriter, message ...interface{}) {
	boom(w, 409, message...)
}

// ResourceGone responds with a 410 Gone error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func ResourceGone(w http.ResponseWriter, message ...interface{}) {
	boom(w, 410, message...)
}

// LengthRequired responds with a 411 Length Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func LengthRequired(w http.ResponseWriter, message ...interface{}) {
	boom(w, 411, message...)
}

// PreconditionFailed responds with a 412 Precondition Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func PreconditionFailed(w http.ResponseWriter, message ...interface{}) {
	boom(w, 412, message...)
}

// EntityTooLarge responds with a 413 Request Entity Too Large error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func EntityTooLarge(w http.ResponseWriter, message ...interface{}) {
	boom(w, 413, message...)
}

// URITooLong responds with a 414 Request-URI Too Large error
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func URITooLong(w http.ResponseWriter, message ...interface{}) {
	boom(w, 414, message...)
}

// UnsupportedMediaType responds with a 415 Unsupported Media Type error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func UnsupportedMediaType(w http.ResponseWriter, message ...interface{}) {
	boom(w, 415, message...)
}

// RangeNotSatisfiable responds with a 416 Requested Range Not Satisfiable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func RangeNotSatisfiable(w http.ResponseWriter, message ...interface{}) {
	boom(w, 416, message...)
}

// ExpectationFailed responds with a 417 Expectation Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func ExpectationFailed(w http.ResponseWriter, message ...interface{}) {
	boom(w, 417, message...)
}

// Teapot responds with a 418 I'm a Teapot error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func Teapot(w http.ResponseWriter, message ...interface{}) {
	boom(w, 418, message...)
}

// BadData responds with a 422 Unprocessable Entity error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func BadData(w http.ResponseWriter, message ...interface{}) {
	boom(w, 422, message...)
}

// Locked responds with a 423 Locked error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func Locked(w http.ResponseWriter, message ...interface{}) {
	boom(w, 423, message...)
}

// PreconditionRequired responds with a 428 Precondition Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func PreconditionRequired(w http.ResponseWriter, message ...interface{}) {
	boom(w, 428, message...)
}

// TooManyRequests responds with a 429 Too Many Requests error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TooManyRequests(w http.ResponseWriter, message ...interface{}) {
	boom(w, 429, message...)
}

// Illegal responds with a 451 Unavailable For Legal Reasons error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func Illegal(w http.ResponseWriter, message ...interface{}) {
	boom(w, 451, message...)
}

// Internal responds with a 500 Internal Server Error error.
// Takes an optional message of either type string or type error
// but, will always return a generic message in the response body.
// This behaviour protects the developer from accidentally returning
// sensitive data in the response during a panic.
func Internal(w http.ResponseWriter, message ...interface{}) {
	boom(w, 500)
}

// BadImplementation responds with a 500 Internal Server Error error. Alias for boom.Internal.
// Takes an optional message of either type string or type error
// but, will always return a generic message in the response body.
// This behaviour protects the developer from accidentally returning
// sensitive data in the response during a panic.
func BadImplementation(w http.ResponseWriter, message ...interface{}) {
	boom(w, 500)
}

// NotImplemented responds with a 501 Not Implemented error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func NotImplemented(w http.ResponseWriter, message ...interface{}) {
	boom(w, 501, message...)
}

// BadGateway responds with a 502 Bad Gateway error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func BadGateway(w http.ResponseWriter, message ...interface{}) {
	boom(w, 502, message...)
}

// ServerUnavailable .eturns a 503 Service Unavailable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func ServerUnavailable(w http.ResponseWriter, message ...interface{}) {
	boom(w, 503, message...)
}

// GatewayTimeout responds with a 504 Gateway Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func GatewayTimeout(w http.ResponseWriter, message ...interface{}) {
	boom(w, 504, message...)
}
