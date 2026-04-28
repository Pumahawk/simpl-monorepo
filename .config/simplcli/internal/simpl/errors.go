package simpl

import (
	"fmt"
	"net/http"
)

const (
	Generic TypeErr = iota
	NotFound
	Conflict
)

type ApiError struct {
	Response *http.Response
	Type     TypeErr
	Err      error
}

func CheckApiError(r *http.Response) *ApiError {
	var typeerr TypeErr = Generic
	var message error = nil
	if r.StatusCode == 404 {
		typeerr = NotFound
		message = fmt.Errorf("status code %d, body=%q", r.StatusCode, bodys(r.Body))
	} else if r.StatusCode == 409 {
		typeerr = Conflict
		message = fmt.Errorf("status code %d, body=%q", r.StatusCode, bodys(r.Body))
	} else if r.StatusCode < 200 || r.StatusCode >= 300 {
		message = fmt.Errorf("status code %d, body=%q", r.StatusCode, bodys(r.Body))
	} else if r.StatusCode >= 500 {
		message = fmt.Errorf("status code=%d body=%q", r.StatusCode, bodys(r.Body))
	}
	if message != nil {
		return &ApiError{
			Response: r,
			Type:     typeerr,
			Err:      message,
		}
	} else {
		return nil
	}
}

func (a *ApiError) Error() string {
	return a.Err.Error()
}

func (a *ApiError) Unwrap() error {
	return a.Err
}

type TypeErr int
