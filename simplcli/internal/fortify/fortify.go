package fortify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Force a pause of fortify requests to avoid status_code=429
var concurrentHttpRequest = make(chan any)

func init() {
	go func() {
		for {
			concurrentHttpRequest <- 0
			<-time.After(5 * time.Second)
		}
	}()
}

type AuthData struct {
	Token string
	User  string
}

type TokenFunc = func() (*AuthData, error)

type Client struct {
	BaseUrl     string
	TokenFunc   TokenFunc
	token       *tokenDto
	tokenExpire time.Time
	tokenSync   sync.Mutex
}

type tokenDto struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (c *Client) newRequest(method string, url string, body any) (*http.Request, error) {
	bj := &bytes.Buffer{}
	err := json.NewEncoder(bj).Encode(body)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest(method, url, bj)
	if err != nil {
		return nil, err
	}
	return r, err
}

func (c *Client) doRequest(r *http.Request, responseBody any) (*http.Response, error) {

	// forify doesnt accept muplice requests. Force a pause
	<-concurrentHttpRequest

	token, err := c.tokenize()
	if err != nil {
		return nil, fmt.Errorf("new request tokenization : %w", err)
	}
	r.Header.Add("Authorization", token)

	cl := http.Client{}
	r.Header.Add("Content-Type", "application/json")
	res, err := cl.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("status_code=%d", res.StatusCode)
	}

	if responseBody != nil {
		err = json.NewDecoder(res.Body).Decode(responseBody)
		if err != nil {
			return res, fmt.Errorf("unmashal body: %w", err)
		}
	}
	return res, nil
}

func (c *Client) tokenize() (string, error) {
	c.tokenSync.Lock()
	defer c.tokenSync.Unlock()

	if c.TokenFunc == nil {
		return "", fmt.Errorf("not found token function for fortify client")
	}

	if c.token != nil && time.Now().Before(c.tokenExpire) {
		return "Bearer " + c.token.AccessToken, nil
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

	c.token = tk
	c.tokenExpire = time.Now().Add(time.Second * time.Duration(tk.ExpiresIn))

	return "Bearer " + tk.AccessToken, nil
}

func (c *Client) MarkFalsePositive(releaseId string) error {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v3/releases/", url.PathEscape(releaseId), "/vulnerabilities/bulk-edit")
	if err != nil {
		return fmt.Errorf("url generation: %w", err)
	}
	body := &MarkFalsePositive{
		DeveloperStatus:           "Will Not Fix",
		AuditorStatus:             "Not an Issue",
		IncludeAllVulnerabilities: true,
	}

	req, err := c.newRequest("POST", rawUrl, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	return nil
}
