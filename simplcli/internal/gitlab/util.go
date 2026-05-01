package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

func searchQuery(values url.Values, search any) {
	v := reflect.ValueOf(search)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("invalid kind, expected struct")
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			name := t.Field(i).Tag.Get("search")
			value := v.Field(i).Interface()
			values.Add(name, fmt.Sprintf("%v", value))
		}
	}
}

func searchApi(c *Client, uri string, search any, response any) error {
	if search != nil {
		u, _ := url.Parse(uri)
		q := u.Query()
		searchQuery(q, search)
		u.RawQuery = q.Encode()
		uri = u.String()
	}
	r, err := c.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	httpCl := http.Client{}
	res, err := httpCl.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("status code %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return err
	}
	return nil
}

func request(c *Client, method string, rawUrl string, body any, response any) (*http.Response, error) {
	//Prepare request
	var rqb io.Reader
	if body != nil {
		bs, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rqb = bytes.NewBuffer(bs)
	}
	rq, err := c.NewRequest(method, rawUrl, rqb)
	if err != nil {
		return nil, err
	}

	// Make http request
	client := http.Client{}
	rs, err := client.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()

	// Check http status code for response
	if rs.StatusCode < 200 || rs.StatusCode >= 300 {
		return rs, fmt.Errorf("status code %d", rs.StatusCode)
	}

	// Prepare response body
	if err := json.NewDecoder(rs.Body).Decode(response); err != nil {
		return rs, err
	}

	return rs, nil
}
