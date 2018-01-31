package boom

import (
	"encoding/json"
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
