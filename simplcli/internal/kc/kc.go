package kc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AuthFunc func() (*AuthInfo, error)

type AuthInfo struct {
	Realm     string
	ClientId  string
	Username  string
	Passaword string
}

type tokenInfo struct {
	token  *TokenizeResponseDto
	expire time.Time
}

type Client struct {
	BaseUrl  string
	AuthFunc AuthFunc
	token    *tokenInfo
}

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

	rawUrl, err := url.JoinPath(c.BaseUrl, "realms", url.PathEscape(auth.Realm), "/protocol/openid-connect/token")
	if err != nil {
		return nil, err
	}

	r, err := http.PostForm(rawUrl, url.Values{
		"client_id":  {auth.ClientId},
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

func (c *Client) Realms() ([]RealmsItemResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/admin/realms")
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("GET", rawUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resb := make([]RealmsItemResponseDto, 0)
	if err := json.NewDecoder(res.Body).Decode(&resb); err != nil {
		return nil, err
	}
	return resb, nil
}

func (c *Client) Realm(realm string) (*RealmResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "realms", url.PathEscape(realm))
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("GET", rawUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resb := &RealmResponseDto{}
	if err := json.NewDecoder(res.Body).Decode(resb); err != nil {
		return nil, err
	}
	return resb, nil
}

type RealmExportOpt struct {
	IncludeClients       bool
	ExportGroupsAndRoles bool
}

func (c *Client) RealmExport(realm string, opt *RealmExportOpt) ([]byte, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/admin/realms", url.PathEscape(realm), "partial-export")
	if err != nil {
		return nil, err
	}

	if opt != nil {
		if u, err := url.Parse(rawUrl); err != nil {
			panic(err)
		} else {
			q := u.Query()
			q.Add("exportClients", strconv.FormatBool(opt.IncludeClients))
			q.Add("exportGroupsAndRoles", strconv.FormatBool(opt.ExportGroupsAndRoles))
			u.RawQuery = q.Encode()
			rawUrl = u.String()
		}
	}

	req, err := c.newRequest("POST", rawUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resb := &bytes.Buffer{}
	if _, err := io.Copy(resb, res.Body); err != nil {
		return nil, err
	}
	return resb.Bytes(), nil
}

func (c *Client) RealmImport(content []byte) error {
	rawUrl, err := url.JoinPath(c.BaseUrl, "admin/realms")
	if err != nil {
		return err
	}

	req, err := c.newRequest("POST", rawUrl, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resb := &bytes.Buffer{}
	if _, err := io.Copy(resb, res.Body); err != nil {
		return err
	}
	return nil
}

func (c *Client) RealmDelete(realm string) error {
	rawUrl, err := url.JoinPath(c.BaseUrl, "admin/realms", url.PathEscape(realm))
	if err != nil {
		return err
	}

	req, err := c.newRequest("DELETE", rawUrl, nil)
	if err != nil {
		return err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
