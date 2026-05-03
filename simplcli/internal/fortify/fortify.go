package fortify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type AuthData struct {
	Token string
	User  string
}

type TokenFunc = func() (AuthData, error)

type Client struct {
	BaseUrl   string
	TokenFunc TokenFunc
	token     tokenDto
}

type tokenDto struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (c *Client) newRequest(method string, url string, body io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	token, err := c.tokenize()
	if err != nil {
		return nil, fmt.Errorf("new request tokenization : %w", err)
	}
	r.Header.Add("Authorization", token)
	return r, err
}

func (c *Client) tokenize() (string, error) {
	if c.TokenFunc == nil {
		return "", fmt.Errorf("not found token function for fortify client")
	}

	auth, err := c.TokenFunc()
	if err != nil {
		return "", fmt.Errorf("error get get auth info fortify: %w", err)
	}

	v := make(url.Values)
	v.Add("grant_type", "password")
	v.Add("scope", "api-tenant")
	v.Add("username", auth.User)
	v.Add("password", auth.Token)
	u, err := url.JoinPath(c.BaseUrl, "/oauth/token")
	if err != nil {
		panic(err)
	}
	res, err := http.PostForm(u, v)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", fmt.Errorf("tokenize status_code=%d", res.StatusCode)
	}

	tk := &tokenDto{}
	err = json.NewDecoder(res.Body).Decode(tk)
	if err != nil {
		return "", fmt.Errorf("token unmashal: %w", err)
	}

	return "Bearer " + tk.AccessToken, nil
}
