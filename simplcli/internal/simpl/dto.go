package simpl

type EchoResponseDto struct {
	Username          string `json:"username"`
	Email             string `json:"email"`
	ConnectionStatus  string `json:"connectionStatus"`
	MtlsStatus        string `json:"mtlsStatus"`
	Id                string `json:"id"`
	Organization      string `json:"organization"`
	CreationTimestamp string `json:"creationTimestamp"`
	UpdateTimestamp   string `json:"updateTimestamp"`
	CredentialId      string `json:"credentialId"`
	ExpiryDate        string `json:"expiryDate"`
	// "userIdentityAttributes": [],
	// "identityAttributes": []
}
