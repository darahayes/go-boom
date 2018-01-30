package boom

import (
	"encoding/json"
	"net/http"
)

type boomErr struct {
	ErrorType  string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func boom(w http.ResponseWriter, statusCode int, msg ...interface{}) {

	errorType := codes[statusCode]
	var message string

	if l := len(msg); l == 0 {
		// no args were provided so set message to errorType
		message = errorType
	} else {
		switch msg[0].(type) {
		case string:
			message = msg[0].(string)
		case error:
			message = msg[0].(error).Error()
		default:
			message = errorType
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(boomErr{
		errorType,
		message,
		statusCode,
	})

	if err != nil {
		panic(err)
	}
}

func BadRequest(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 400, msg...)
}

func Unathorized(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 401, msg...)
}

func PaymentRequired(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 402, msg...)
}

func Forbidden(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 403, msg...)
}

func NotFound(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 404, msg...)
}

func MethodNotAllowed(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 405, msg...)
}

func NotAcceptable(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 406, msg...)
}

func ProxyAuthRequired(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 407, msg...)
}

func ClientTimeout(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 408, msg...)
}

func Conflict(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 409, msg...)
}

func ResourceGone(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 410, msg...)
}

func LengthRequired(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 411, msg...)
}

func PreconditionFailed(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 412, msg...)
}

func EntityTooLarge(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 413, msg...)
}

func URITooLong(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 414, msg...)
}

func UnsupportedMediaType(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 415, msg...)
}

func RangeNotSatisfiable(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 416, msg...)
}

func ExpectationFailed(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 417, msg...)
}

func Teapot(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 418, msg...)
}

func BadData(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 422, msg...)
}

func PreconditionRequired(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 428, msg...)
}

func TooManyRequests(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 429, msg...)
}

func Illegal(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 451, msg...)
}

// 5xx errors - make sure no error message is returned to client
func BadImplementation(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 500)
}

func NotImplemented(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 501)
}

func BadGateway(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 502)
}

func ServerUnavailable(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 503)
}

func GatewayTimeout(w http.ResponseWriter, msg ...interface{}) {
	boom(w, 504)
}
