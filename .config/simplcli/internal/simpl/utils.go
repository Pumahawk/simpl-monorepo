package simpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

// Generic function for check healt
func microCheck(url string) (bool, error) {
	r, err := http.Get(url + "/actuator/health")
	if err != nil {
		return false, err
	}
	defer r.Body.Close()

	rj := &healtResponseDto{}
	if err := json.NewDecoder(r.Body).Decode(rj); err != nil {
		return false, err
	}

	return rj.Status == "UP", nil
}

func jsonR(body any) io.Reader {
	if body == nil {
		return nil
	}
	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(b)
}

func request(method, url string, bodyRequest any, bodyResponse any) (*http.Response, error) {
	// Prepare request
	r, err := http.NewRequest(method, url, jsonR(bodyRequest))
	if err != nil {
		// Unexpected exception [bug]
		panic(err)
	}
	if bodyRequest != nil {
		r.Header.Add("Content-Type", "application/json")
	}

	// Do request
	hc := http.Client{}
	rh, err := hc.Do(r)
	if err != nil {
		return rh, err
	}
	defer rh.Body.Close()

	// Check response code
	if apiErr := CheckApiError(rh); apiErr != nil {
		return rh, apiErr
	}

	// Extract response body
	if bodyResponse != nil {
		if err := jsonDec(rh.Body, bodyResponse); err != nil {
			return nil, err
		}
	}

	return rh, nil
}

func jsonDec(r io.Reader, v any) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}
	return nil
}

func bodys(r io.Reader) string {
	b, err := io.ReadAll(r)
	if err != nil {
		return fmt.Sprintf("err=%q", err.Error())
	}
	return string(b)
}

func searchApi(page, size int, rawUrl string, filters any, bodyResponse any) (*http.Response, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("size", strconv.Itoa(size))
	if filters != nil {
		queryFilters(q, filters)
	}
	u.RawQuery = q.Encode()

	return request("GET", u.String(), nil, bodyResponse)
}

func queryFilters(values url.Values, filters any) {
	v := reflect.ValueOf(filters)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	t := v.Type()
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("unsupported type filters %T", filters))
	}
	for i := range v.NumField() {
		fv := v.Field(i)
		ft := t.Field(i)
		name := ft.Tag.Get("search")
		if name == "" {
			name = ft.Name
		}
		value := fv.String()
		if value != "" {
			values.Add(name, value)
		}
	}
}
