package simpl

import (
	"net/url"
	"strconv"
)

type Client struct {
	authenticationProviderUrl     string
	identityProviderUrl           string
	onboardingUrl                 string
	securityAttributesProviderUrl string
	tier1GatewayUrl               string
	tier2GatewayUrl               string
	usersRolesUrl                 string
}

func (c *Client) KeypairsActive() (bool, error) {
	url := c.authenticationProviderUrl + "/tier1/v2/keypairs/active"
	if _, err := request("HEAD", url, nil, nil); err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) GenerateKeypair(name string) (*GenerateKeypairResponseDto, error) {
	url := c.authenticationProviderUrl + "/tier1/v2/keypairs"
	body := &GenerateKeypairRequestDto{Name: name}
	rdto := &GenerateKeypairResponseDto{}
	if _, err := request("POST", url, body, rdto); err != nil {
		return nil, err
	}
	return rdto, nil
}

func (c *Client) Keypairs(page, size int, search *KeypairsSearch) (*KeypairsResponseDto, error) {
	u, err := url.Parse(c.authenticationProviderUrl + "/tier1/v2/keypairs")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("size", strconv.Itoa(size))
	if search != nil {
		search.Query(q)
	}
	u.RawQuery = q.Encode()

	r := &KeypairsResponseDto{}
	if _, err := request("GET", u.String(), nil, r); err != nil {
		return nil, err
	}
	return r, nil
}
