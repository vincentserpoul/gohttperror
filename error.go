package gohttperror

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/vincentserpoul/gohttpmw"
)

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error
}

// Render rendering the error
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	gohttpmw.SetRequestError(r, e.Err)
	return nil
}

// ErrBadRequest when supplied data is not correct
func ErrBadRequest(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
	}
}

// ErrInternal when there is a server side issue
func ErrInternal(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Error rendering response.",
	}
}

// ErrUnauthorized when there is a server side issue
func ErrUnauthorized(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     "Not authorized",
	}
}

// ErrForbidden when there is a server side issue
func ErrForbidden(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		StatusText:     "Forbidden",
	}
}

// ErrNotFound is the wrapped error for not found resources
var ErrNotFound = &ErrResponse{
	HTTPStatusCode: http.StatusNotFound,
	StatusText:     "Resource not found",
}

// ErrNotImplemented is the wrapped error for not implemented methods
var ErrNotImplemented = &ErrResponse{
	HTTPStatusCode: http.StatusNotImplemented,
	StatusText:     "Method not implemented",
}
