package simpl

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
