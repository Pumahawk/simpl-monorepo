package simpl

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
