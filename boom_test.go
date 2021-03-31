package boom

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func renderBoom(w http.ResponseWriter, statusCode int, args ...interface{}) {
	boomed := Boomify(statusCode, args...)
	Render(w, boomed)
}

func TestBoomDefault(t *testing.T) {
	rr := httptest.NewRecorder()

	var boomResponse Err
	statusCode := 404
	errorType := codes[statusCode]
	message := errorType

	renderBoom(rr, statusCode)

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

func TestBoomCustomMessage(t *testing.T) {
	rr := httptest.NewRecorder()

	var boomResponse Err
	statusCode := 404
	errorType := codes[statusCode]
	message := "This is a test message"

	renderBoom(rr, statusCode, message)

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

func TestBoomCustomErrorMessage(t *testing.T) {
	rr := httptest.NewRecorder()

	var boomResponse Err
	statusCode := 404
	errorType := codes[statusCode]
	message := errors.New("This is a message of type error")

	renderBoom(rr, statusCode, message)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != message.Error() || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}

func TestBoomHandlesMultipleMessagesWithNoFailure(t *testing.T) {
	rr := httptest.NewRecorder()

	var boomResponse Err
	statusCode := 404
	// errorType := codes[statusCode]
	m1 := "This is a test message"
	m2 := "This is another test message"
	m3 := "This is a third test message"

	renderBoom(rr, statusCode, m1, m2, m3)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}
