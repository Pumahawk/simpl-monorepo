package simpl

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
	url := c.authenticationProviderUrl + "/tier1/v2/keypairs"
	r := &KeypairsResponseDto{}
	if _, err := searchApi(page, size, url, search, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) Participants(page, size int, search *ParticipantsSearch) (*ParticipantsResponseDto, error) {
	url := c.identityProviderUrl + "/tier1/v2/participants"
	r := &ParticipantsResponseDto{}
	if _, err := searchApi(page, size, url, search, r); err != nil {
		return nil, err
	}
	return r, nil
}
