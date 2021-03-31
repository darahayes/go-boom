package boom

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

// BadRequest responds with a 400 Bad Request error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestBadRequest(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 400
	errorType := "Bad Request"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderBadRequest(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// Unathorized responds with a 401 Unauthorized error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestUnathorized(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 401
	errorType := "Unauthorized"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderUnathorized(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// PaymentRequired responds with a 402 Payment Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestPaymentRequired(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 402
	errorType := "Payment Required"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderPaymentRequired(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// Forbidden responds with a 403 Forbidden error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestForbidden(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 403
	errorType := "Forbidden"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderForbidden(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// NotFound responds with a 404 Not Found error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestNotFound(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 404
	errorType := "Not Found"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderNotFound(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// MethodNotAllowed responds with a 405 Method Not Allowed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestMethodNotAllowed(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 405
	errorType := "Method Not Allowed"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderMethodNotAllowed(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// NotAcceptable responds with a 406 Not Acceptable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestNotAcceptable(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 406
	errorType := "Not Acceptable"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderNotAcceptable(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// ProxyAuthRequired responds with a 407 Proxy Authentication Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestProxyAuthRequired(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 407
	errorType := "Proxy Authentication Required"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderProxyAuthRequired(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// ClientTimeout responds with a 408 Request Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestClientTimeout(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 408
	errorType := "Request Time-out"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderRequestTimeout(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// Conflict responds with a 409 Conflict error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestConflict(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 409
	errorType := "Conflict"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderConflict(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// ResourceGone responds with a 410 Gone error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestResourceGone(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 410
	errorType := "Gone"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderGone(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// LengthRequired responds with a 411 Length Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestLengthRequired(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 411
	errorType := "Length Required"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderLengthRequired(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// PreconditionFailed responds with a 412 Precondition Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestPreconditionFailed(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 412
	errorType := "Precondition Failed"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderPreconditionFailed(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// EntityTooLarge responds with a 413 Request Entity Too Large error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestEntityTooLarge(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 413
	errorType := "Request Entity Too Large"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderEntityTooLarge(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// URITooLong responds with a 414 Request-URI Too Large error
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestURITooLong(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 414
	errorType := "Request-URI Too Large"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderURITooLong(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// UnsupportedMediaType responds with a 415 Unsupported Media Type error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestUnsupportedMediaType(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 415
	errorType := "Unsupported Media Type"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderUnsupportedMediaType(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// RangeNotSatisfiable responds with a 416 Requested Range Not Satisfiable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestRangeNotSatisfiable(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 416
	errorType := "Requested Range Not Satisfiable"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderRangeNotSatisfiable(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// ExpectationFailed responds with a 417 Expectation Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestExpectationFailed(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 417
	errorType := "Expectation Failed"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderExpectationFailed(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// Teapot responds with a 418 I'm a Teapot error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestTeapot(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 418
	errorType := "I'm a teapot"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderTeapot(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// BadData responds with a 422 Unprocessable Entity error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestBadData(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 422
	errorType := "Unprocessable Entity"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderUnprocessableEntity(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// Locked responds with a 423 Locked error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestLocked(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 423
	errorType := "Locked"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderLocked(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// PreconditionRequired responds with a 428 Precondition Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestPreconditionRequired(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 428
	errorType := "Precondition Required"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderPreconditionRequired(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// TooManyRequests responds with a 429 Too Many Requests error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestTooManyRequests(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 429
	errorType := "Too Many Requests"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderTooManyRequests(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// Illegal responds with a 451 Unavailable For Legal Reasons error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestIllegal(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 451
	errorType := "Unavailable For Legal Reasons"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderIllegal(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// Internal responds with a 500 Internal Server Error error.
// Takes an optional message of either type string or type error
// but, will always return a generic message in the response body.
// This behaviour protects the developer from accidentally returning
// sensitive data in the response during a panic.
func TestInternal(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 500
	errorType := "Internal Server Error"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderInternal(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	// 500 is special. It should never return a custom message. Should always be same as errorType
	if boomResponse.ErrorType != errorType || boomResponse.Message != errorType || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// BadImplementation responds with a 500 Internal Server Error error. Alias for boom.Internal.
// Takes an optional message of either type string or type error
// but, will always return a generic message in the response body.
// This behaviour protects the developer from accidentally returning
// sensitive data in the response during a panic.
func TestBadImplementation(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 500
	errorType := "Internal Server Error"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderBadImplementation(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	// 500 is special. It should never return a custom message. Should always be same as errorType
	if boomResponse.ErrorType != errorType || boomResponse.Message != errorType || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// NotImplemented responds with a 501 Not Implemented error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestNotImplemented(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 501
	errorType := "Not Implemented"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderNotImplemented(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// BadGateway responds with a 502 Bad Gateway error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestBadGateway(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 502
	errorType := "Bad Gateway"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderBadGateway(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// ServerUnavailable .eturns a 503 Service Unavailable error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestServerUnavailable(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 503
	errorType := "Service Unavailable"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderServiceUnavailable(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

// GatewayTimeout responds with a 504 Gateway Time-out error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func TestGatewayTimeout(t *testing.T) {
	rr := httptest.NewRecorder()

	statusCode := 504
	errorType := "Gateway Time-out"
	message := "This is a custom messsage"

	var boomResponse Err

	RenderGatewayTimeout(rr, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}
