package simpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var DefaultEndpoint *SingleAccessEndpoints = &SingleAccessEndpoints{
	BaseUrl: "http://localhost:8100",
}

type TokenFunc func() (string, error)

type Client struct {
	Endpoints Endpoints
	TokenFunc TokenFunc
}

func (c *Client) edp() Endpoints {
	if c.Endpoints != nil {
		return c.Endpoints
	} else {
		return DefaultEndpoint
	}
}

type Endpoints interface {
	Keycloak() string
	AuthenticationProvider() string
	IdentityProvider() string
}

type SingleAccessEndpoints struct {
	BaseUrl                   string
	KeycloakUrl               string
	AuthenticationProviderUrl string
	IdentityProviderUrl       string
}

func (s *SingleAccessEndpoints) ordef(val, def string) string {
	if val != "" {
		return val
	} else {
		return def
	}
}

func (s *SingleAccessEndpoints) Keycloak() string {
	return s.BaseUrl + s.ordef(s.KeycloakUrl, "/auth")
}

func (s *SingleAccessEndpoints) AuthenticationProvider() string {
	return s.BaseUrl + s.ordef(s.AuthenticationProviderUrl, "/authApi")
}

func (s *SingleAccessEndpoints) IdentityProvider() string {
	return s.BaseUrl + s.ordef(s.IdentityProviderUrl, "/identityApi")
}

func (c *Client) newRequest(method, url string, body any) (*http.Request, error) {
	var r io.Reader
	if body != nil {
		bf := &bytes.Buffer{}
		if err := json.NewEncoder(bf).Encode(body); err != nil {
			return nil, err
		}
		r = bf
	}
	rq, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}

	if r != nil {
		rq.Header.Add("Content-Type", "application/json")
	}
	return rq, nil
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	token, err := c.TokenFunc()
	if err != nil {
		return nil, fmt.Errorf("tokenize error: %w", err)
	}

	req.Header.Add("authorization", "Bearer "+token)

	cl := &http.Client{}
	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return res, fmt.Errorf("simpl api httpcode=%d", res.StatusCode)
	}

	return res, nil
}
