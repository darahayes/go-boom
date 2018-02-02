package boom

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
)

func TestBoomDefault(t *testing.T) {
	rr := httptest.NewRecorder()

	var boomResponse boomErr
	statusCode := 404
	errorType := codes[statusCode]
	message := errorType

	boom(rr, statusCode)

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

	var boomResponse boomErr
	statusCode := 404
	errorType := codes[statusCode]
	message := "This is a test message"

	boom(rr, statusCode, message)

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

	var boomResponse boomErr
	statusCode := 404
	errorType := codes[statusCode]
	message := errors.New("This is a message of type error")

	boom(rr, statusCode, message)

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

	var boomResponse boomErr
	statusCode := 404
	errorType := codes[statusCode]
	m1 := "This is a test message"
	m2 := "This is another test message"
	m3 := "This is a third test message"

	boom(rr, statusCode, m1, m2, m3)

	if err := json.Unmarshal(rr.Body.Bytes(), &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	if boomResponse.ErrorType != errorType || boomResponse.Message != m1 || boomResponse.StatusCode != statusCode {
		t.Fail()
	}

	if rr.Code != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, statusCode)
	}
}
