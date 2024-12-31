package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Load() int64 {
	// load app api keys, passwords, etc

	return 1
}

func LoadGlobals() int64 {
	return 2
}

type HTTPError struct {
	Code    int
	Err     error
	Message string
}

func (e *HTTPError) Error() string {
	return e.Err.Error() + ": " + e.Message
}

func HTTPErrorHandler(err error, c echo.Context) {
	var _code int = http.StatusInternalServerError
	var _msg string = http.StatusText(_code)

	switch err.(type) {
	case *echo.HTTPError:
		he, _ := err.(*echo.HTTPError)
		_code = he.Code
		_msg = fmt.Sprintf("%v", he.Message)

	case *HTTPError:
		he, _ := err.(*HTTPError)
		_code = he.Code
		_msg = he.Message

	default:
		he, _ := err.(*HTTPError)
		_code = he.Code
		_msg = he.Message

	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			c.NoContent(_code)
		} else {
			c.String(_code, _msg)
		}
	}
	log.Println(err.Error())
}
