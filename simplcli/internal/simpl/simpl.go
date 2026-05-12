package simpl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type TokenizeResponseDto struct {
	AccessToken                 string `json:"access_token"`
	ExpiresIn                   int    `json:"expires_in"`
	RefreshExpiresIn            int    `json:"refresh_expires_in"`
	RefreshToken                string `json:"refresh_token"`
	TokenType                   string `json:"token_type"`
	SessionState                string `json:"session_state"`
	Scope                       string `json:"scope"`
	NotBeforePolicyBeforePolicy int    `json:"not-before-policy"`
}

func (c *Client) Tokenize() (*TokenizeResponseDto, error) {
	if c.AuthFunc == nil {
		return nil, fmt.Errorf("authentication function not defined")
	}

	if c.token != nil && time.Now().Before(c.token.expire) {
		return c.token.token, nil
	}

	auth, err := c.AuthFunc()
	if err != nil {
		return nil, fmt.Errorf("authentication func problem: %w", err)
	}

	rawUrl, err := url.JoinPath(c.edp().Keycloak(), "realms", url.PathEscape(auth.Realm), "/protocol/openid-connect/token")
	if err != nil {
		return nil, err
	}

	r, err := http.PostForm(rawUrl, url.Values{
		"client_id":  {"frontend-cli"},
		"username":   {auth.Username},
		"password":   {auth.Passaword},
		"grant_type": {"password"},
	})
	if err != nil {
		return nil, fmt.Errorf("postform token: %w", err)
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		return nil, fmt.Errorf("post token httpcode=%d", r.StatusCode)
	}

	defer r.Body.Close()

	res := &TokenizeResponseDto{}
	err = json.NewDecoder(r.Body).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("read token response body: %w", err)
	}

	return res, nil
}

func (c *Client) Echo() (*EchoResponseDto, error) {
	rawUrl, err := url.JoinPath(c.edp().AuthenticationProvider(), "/tier1/v2/echo")
	if err != nil {
		return nil, err
	}
	rq, err := c.newRequest("GET", rawUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(rq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rb := &EchoResponseDto{}
	err = json.NewDecoder(res.Body).Decode(rb)
	if err != nil {
		return nil, err
	}

	return rb, nil
}
