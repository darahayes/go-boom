package boom

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoverHandlerPanic(t *testing.T) {

	var boomResponse Err

	fn := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		panic("This is a purposeful panic!")
	})

	handler := RecoverHandler(fn)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Errorf("Request to server failed %v", err)
	}

	if res != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	expectedErrorType := "Internal Server Error"
	expectedMessage := expectedErrorType
	expectedCode := 500

	if boomResponse.ErrorType != expectedErrorType || boomResponse.Message != expectedMessage || boomResponse.StatusCode != expectedCode {
		t.Fail()
	}

	if res.StatusCode != expectedCode {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, expectedCode)
	}
}

func TestRecoverHandlerNoPanic(t *testing.T) {

	fn := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "hello world!")
	})

	handler := RecoverHandler(fn)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Errorf("Request to server failed %v", err)
	}

	if res != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)

	if string(body) != "hello world!" || res.StatusCode != 200 {
		t.Fail()
	}
}

func TestNotFoundHandler(t *testing.T) {

	var boomResponse Err

	handler := http.HandlerFunc(NotFoundHandler)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Errorf("Request to server failed %v", err)
	}

	if res != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &boomResponse); err != nil {
		t.Errorf("response body was not valid JSON: %v", err)
	}

	expectedErrorType := "Not Found"
	expectedMessage := expectedErrorType
	expectedCode := 404

	if boomResponse.ErrorType != expectedErrorType || boomResponse.Message != expectedMessage || boomResponse.StatusCode != expectedCode {
		t.Fail()
	}

	if res.StatusCode != expectedCode {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, expectedCode)
	}
}
