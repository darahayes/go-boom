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

func boom(w http.ResponseWriter, statusCode int, args ...interface{}) {

	errorType := codes[statusCode]
	message := errorType // should be same as errorType by default

	// determine if an error or string arg was passed in
	// set the message accordingly
	if l := len(args); l != 0 {
		switch args[0].(type) {
		case string:
			message = args[0].(string)
		case error:
			message = args[0].(error).Error()
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
