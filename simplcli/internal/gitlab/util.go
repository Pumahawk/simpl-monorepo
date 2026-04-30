package gitlab

import (
	"encoding/json"
	"fmt"
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

func searchApi(_ *Client, uri string, search any, response any) error {
	if search != nil {
		u, _ := url.Parse(uri)
		q := u.Query()
		searchQuery(q, search)
		u.RawQuery = q.Encode()
		uri = u.String()
	}
	r, err := http.NewRequest("GET", uri, nil)
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
