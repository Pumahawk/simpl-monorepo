package simpl

import (
	"encoding/json"
	"net/url"
)

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
