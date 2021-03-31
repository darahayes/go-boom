package boom

import (
	"encoding/json"
	"net/http"
)

var codes = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Moved Temporarily",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	307: "Temporary Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Time-out",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-URI Too Large",
	415: "Unsupported Media Type",
	416: "Requested Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Unordered Collection",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-out",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	509: "Bandwidth Limit Exceeded",
	510: "Not Extended",
	511: "Network Authentication Required",
}

// Err holds information about the error to be returned to the client
// during a renderBoom function call or elsewhere
type Err struct {
	ErrorType  string             `json:"error"`
	Message    string             `json:"message"`
	StatusCode int                `json:"statusCode"`
	Data       *map[string]string `json:"data,omitempty"`
}

func Boomify(statusCode int, args ...interface{}) Err {
	errorType := codes[statusCode]
	message := errorType // should be same as errorType by default

	// determine if an error or string arg was passed in
	// set the message accordingly
	var data *map[string]string
	if l := len(args); l != 0 {
		for _, arg := range args {
			switch arg.(type) {
			case string:
				message = arg.(string)
			case error:
				message = arg.(error).Error()
			case map[string]string:
				dataRaw := arg.(map[string]string)
				data = &dataRaw
			}
		}
	}
	return Err{
		errorType,
		message,
		statusCode,
		data,
	}
}

func Render(w http.ResponseWriter, boomed Err) {
	errString, _ := json.Marshal(boomed)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(boomed.StatusCode)
	w.Write(errString)
}
