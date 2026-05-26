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

func (c *Client) KeypairActive() (*KeyPairActiveDto, error) {
	rawUrl, err := url.JoinPath(c.edp().AuthenticationProvider(), "/tier1/v2/keypairs/active")
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

	rb := &KeyPairActiveDto{}
	err = json.NewDecoder(res.Body).Decode(rb)
	if err != nil {
		return nil, err
	}

	return rb, nil
}

func (c *Client) GenerateKeyPair(name string) (*GenerateKeyPairResponseDto, error) {
	rawUrl, err := url.JoinPath(c.edp().AuthenticationProvider(), "/tier1/v2/keypairs")
	if err != nil {
		return nil, err
	}
	rqb := &GenerateKeyPairRequestDto{name}
	rq, err := c.newRequest("POST", rawUrl, rqb)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(rq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rb := &GenerateKeyPairResponseDto{}
	err = json.NewDecoder(res.Body).Decode(rb)
	if err != nil {
		return nil, err
	}

	return rb, nil
}

func (c *Client) DataspaceUpdate(name, email string) error {
	rawUrl, err := url.JoinPath(c.edp().AuthenticationProvider(), "/tier1/v2/identity/dataspace")
	if err != nil {
		return err
	}
	rqb := &DataspaceUpdateDto{name, email}
	rq, err := c.newRequest("PUT", rawUrl, rqb)
	if err != nil {
		return err
	}

	res, err := c.doRequest(rq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (c *Client) CreateCsr(keypairId, cn string) error {
	rawUrl, err := url.JoinPath(c.edp().AuthenticationProvider(), "/tier1/v2/identity/dataspace")
	if err != nil {
		return err
	}
	rqb := &CsrRequest{
		CommonName:         cn,
		Country:            cn,
		Organization:       cn,
		OrganizationalUnit: cn,
	}
	rq, err := c.newRequest("POST", rawUrl, rqb)
	if err != nil {
		return err
	}

	res, err := c.doRequest(rq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (c *Client) ParticipantCreate(partType, organization string, isAuthority bool) error {
	rawUrl, err := url.JoinPath(c.edp().IdentityProvider(), "/tier1/v2/identity/dataspace")
	if err != nil {
		return err
	}
	rqb := &ParticipantCreateRequestDto{
		Organization:    organization,
		ParticipantType: partType,
		IsAuthority:     isAuthority,
	}
	rq, err := c.newRequest("POST", rawUrl, rqb)
	if err != nil {
		return err
	}

	res, err := c.doRequest(rq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (c *Client) UploadCsr(keypairId, cn string) error {
	rawUrl, err := url.JoinPath(c.edp().IdentityProvider(), "/tier1/v2/identity/dataspace")
	if err != nil {
		return err
	}
	rqb := &CsrRequest{
		CommonName:         cn,
		Country:            cn,
		Organization:       cn,
		OrganizationalUnit: cn,
	}
	rq, err := c.newRequest("POST", rawUrl, rqb)
	if err != nil {
		return err
	}

	res, err := c.doRequest(rq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
