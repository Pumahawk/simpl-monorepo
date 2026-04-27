package simpl

import (
	"encoding/json"
	"net/http"
)

// Default authority endpoints
const authenticationProviderUrlAuthority = "http://localhost:8105"
const identityProviderUrlAuthority = "http://localhost:8103"
const onboardingUrlAuthority = "http://localhost:8104"
const securityAttributesProviderUrlAuthority = "http://localhost:8102"
const tier1GatewayUrlAuthority = "http://localhost:8100"
const tier2GatewayUrlAuthority = "http://localhost:8142"
const usersRolesUrlAuthority = "http://localhost:8101"

// Default participant endpoints
const authenticationProviderUrlParticipant = "http://localhost:8205"
const identityProviderUrlParticipant = "http://localhost:8203"
const onboardingUrlParticipant = "http://localhost:8204"
const securityAttributesProviderUrlParticipant = "http://localhost:8202"
const tier1GatewayUrlParticipant = "http://localhost:8200"
const tier2GatewayUrlParticipant = "http://localhost:8242"
const usersRolesUrlParticipant = "http://localhost:8201"

type Client struct {
	authenticationProviderUrl     string
	identityProviderUrl           string
	onboardingUrl                 string
	securityAttributesProviderUrl string
	tier1GatewayUrl               string
	tier2GatewayUrl               string
	usersRolesUrl                 string
}

// Constructor deefinitions
func NewAuthorityClient() Client {
	return Client{
		authenticationProviderUrl:     authenticationProviderUrlAuthority,
		identityProviderUrl:           identityProviderUrlAuthority,
		onboardingUrl:                 onboardingUrlAuthority,
		securityAttributesProviderUrl: securityAttributesProviderUrlAuthority,
		tier1GatewayUrl:               tier1GatewayUrlAuthority,
		tier2GatewayUrl:               tier2GatewayUrlAuthority,
		usersRolesUrl:                 usersRolesUrlAuthority,
	}
}

func NewParticipantClient() Client {
	return Client{
		authenticationProviderUrl:     authenticationProviderUrlParticipant,
		identityProviderUrl:           identityProviderUrlParticipant,
		onboardingUrl:                 onboardingUrlParticipant,
		securityAttributesProviderUrl: securityAttributesProviderUrlParticipant,
		tier1GatewayUrl:               tier1GatewayUrlParticipant,
		tier2GatewayUrl:               tier2GatewayUrlParticipant,
		usersRolesUrl:                 usersRolesUrlParticipant,
	}
}

// Check healt functions
func (c *Client) CheckAuthenticationProvider() (bool, error) {
	return microCheck(c.authenticationProviderUrl)
}

func (c *Client) CheckIdentityProvider() (bool, error) {
	return microCheck(c.identityProviderUrl)
}

func (c *Client) CheckOnboarding() (bool, error) {
	return microCheck(c.onboardingUrl)
}

func (c *Client) CheckSecurityAttributesProvider() (bool, error) {
	return microCheck(c.securityAttributesProviderUrl)
}

func (c *Client) CheckTier1Gateway() (bool, error) {
	return microCheck(c.tier1GatewayUrl)
}

func (c *Client) CheckTier2Gateway() (bool, error) {
	return microCheck(c.tier2GatewayUrl)
}

func (c *Client) CheckUsersRoles() (bool, error) {
	return microCheck(c.usersRolesUrl)
}

// Generic function for check healt
func microCheck(url string) (bool, error) {
	r, err := http.Get(url + "/actuator/health")
	if err != nil {
		return false, err
	}
	defer r.Body.Close()

	rj := &healtResponseDto{}
	if err := json.NewDecoder(r.Body).Decode(rj); err != nil {
		return false, err
	}

	return rj.Status == "UP", nil
}

// Check health DTO
type healtResponseDto struct {
	Status string
}
