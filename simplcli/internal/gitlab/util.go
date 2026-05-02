package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type searchTags struct {
	Name    string
	Default reflect.Value
}

func getSearchTags(field reflect.StructField) searchTags {
	tagSearch := field.Tag.Get("search")
	tags := strings.Split(tagSearch, ",")
	name := tags[0]
	tagm := make(map[string]string)
	for _, t := range tags[1:] {
		kv := strings.Split(t, ":")
		k := kv[0]
		v := strings.Join(kv[1:], ":")
		tagm[k] = v
	}
	value := reflect.New(field.Type).Elem()
	if def, ok := tagm["def"]; ok {
		switch field.Type.Kind() {
		case reflect.String:
			value = reflect.ValueOf(def)
		case reflect.Int:
			value = strconfOrPanic(def, strconv.Atoi)
		case reflect.Float64:
			value = strconfOrPanic(def, func(def string) (float64, error) { return strconv.ParseFloat(def, 64) })
		case reflect.Float32:
			value = strconfOrPanic(def, func(def string) (float32, error) { v, err := strconv.ParseFloat(def, 32); return float32(v), err })
		}
	}
	return searchTags{
		Name:    name,
		Default: value,
	}
}

func strconfOrPanic[T any](str string, parser func(string) (T, error)) reflect.Value {
	v, err := parser(str)
	if err != nil {
		panic(err)
	}
	return reflect.ValueOf(v)
}

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
		tags := getSearchTags(t.Field(i))
		name := tags.Name
		var valueV reflect.Value
		if v.Field(i).IsZero() {
			valueV = tags.Default
		} else {
			valueV = v.Field(i)
		}
		if !valueV.IsZero() {
			values.Add(name, fmt.Sprintf("%v", valueV.Interface()))
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
