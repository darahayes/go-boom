# Boom

[![Documentation](https://godoc.org/github.com/darahayes/go-boom?status.svg)](http://godoc.org/github.com/darahayes/go-boom)
[![Coverage Status](https://coveralls.io/repos/github/darahayes/go-boom/badge.svg?branch=master)](https://coveralls.io/github/darahayes/go-boom?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/darahayes/go-boom)](https://goreportcard.com/report/github.com/darahayes/go-boom)
[![CircleCI](https://circleci.com/gh/darahayes/go-boom.svg?style=svg)](https://circleci.com/gh/darahayes/go-boom)

**boom** provides a set of functions for returning HTTP errors. Each function responds with a JSON object which includes the following properties:

- `statusCode` - the HTTP status code.
- `error`- the HTTP status message (e.g. 'Bad Request', 'Internal Server Error') derived from `statusCode`.
- `message` - the error message which can be optionally set by the developer.

All **boom** functions take a `http.ResponseWriter` as an argument which means **boom** should be compatible with any Golang http frameworks that also use `http.ResponseWriter`.

To see the full list of **boom** functions, check out the [documentation on godoc.org](https://godoc.org/github.com/darahayes/go-boom)

This library is inspired by the wonderful JavaScript library https://www.npmjs.com/package/boom.

### Install

```bash
go get github.com/darahayes/go-boom
```

### Import

```go
import (
	"github.com/darahayes/go-boom"
)
```

## Example Usage

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/darahayes/go-boom"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	boom.NotFound(w, "Sorry, there's nothing here.")
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
```

With this example, the response from the `/` endpoint would be:

```json
{
  "error": "Not Found",
  "message": "Sorry, there's nothing here.",
  "statusCode": 404
}
```

**boom** also accepts arguments of type `error`, or any `struct` that implements the built in `error` interface. (i.e. has an `Error()` function which returns a `string`)

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("You shall not pass!")
	boom.Unauthorized(w, err)
}
```

and the response:

```json
{
  "error": "Unauthorized",
  "message": "You shall not pass!",
  "statusCode": 401
}
```
It's also possible to provide no error message at all:

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
	boom.BadRequest(w)
}
```

and the response:

```json
{
  "error": "Bad Request",
  "message": "Bad Request",
  "statusCode": 400
}
```

## boom.RecoverHandler

**boom** also comes with a `RecoverHandler` middleware function that can be used to recover from unexpected panics.
It can be used directly or with any router library that deals with the `http.Handler` type. (most of them do!)

`boom.RecoverHandler` does three things:

1. Recovers from unexpected panics.
2. Logs a stack trace in the server logs.
3. Uses `boom.Internal()` to return a generic `500 Internal Server Error`. This prevents accidental leakage of sensitive info in the response.

#### Example usage with plain `http` library

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
	panic("Uh oh, something happened")
}

func main() {
	http.Handle("/", boom.RecoverHandler(http.HandlerFunc(myHandler)))
	http.ListenAndServe(":8080", nil)
}
```

#### Example usage with the `mux` router

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
	panic("Uh oh, something happened")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", myHandler)
	
	router.Use(boom.RecoverHandler)
	
	http.ListenAndServe(":8080", router)
}
```

In both examples above, a stack trace is printed in the server logs and the response is the following:

```
{
  "error": "Internal Server Error",
  "message": "Internal Server Error",
  "statusCode": 500
}
```

## API Methods

To see the full list of API methods check out the [documentation on godoc.org](https://godoc.org/github.com/darahayes/go-boom)
