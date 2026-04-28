package simpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	if rh.StatusCode == 404 {
		return rh, fmt.Errorf("status code %d: %w", rh.StatusCode, NotFound)
	} else if rh.StatusCode < 200 || rh.StatusCode >= 300 {
		return rh, fmt.Errorf("status code %d, body=%q", rh.StatusCode, bodys(rh.Body))
	} else if rh.StatusCode >= 500 {
		return rh, fmt.Errorf("status code=%d body=%q", rh.StatusCode, bodys(rh.Body))
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
