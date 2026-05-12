package simpl

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var DefaultEndpoint *SingleAccessEndpoints = &SingleAccessEndpoints{
	BaseUrl: "http://localhost:8100",
}

type AuthFunc func() (*AuthInfo, error)

type AuthInfo struct {
	Realm     string
	Username  string
	Passaword string
}

type tokenInfo struct {
	token  *TokenizeResponseDto
	expire time.Time
}

type Client struct {
	Endpoints Endpoints
	AuthFunc  AuthFunc
	token     *tokenInfo
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
}

type SingleAccessEndpoints struct {
	BaseUrl string
}

func (s *SingleAccessEndpoints) Keycloak() string {
	return s.BaseUrl + "/auth"
}

func (s *SingleAccessEndpoints) AuthenticationProvider() string {
	return s.BaseUrl + "/authApi"
}

func (c *Client) apiToken() (string, error) {
	if c.token == nil || time.Now().Before(c.token.expire) {
		tk, err := c.Tokenize()
		if err != nil {
			return "", err
		}

		expire := time.Now().Add(time.Duration(tk.ExpiresIn) * time.Second)

		c.token = &tokenInfo{
			token:  tk,
			expire: expire,
		}
	}

	return c.token.token.AccessToken, nil
}

func (c *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	token, err := c.apiToken()
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
