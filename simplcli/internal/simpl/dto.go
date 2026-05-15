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

type KeyPairActiveDto struct {
	Active            bool   `json:"active"`
	CreationTimestamp string `json:"creationTimestamp"`
	Csr               string `json:"csr"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	PrivateKey        string `json:"privateKey"`
	PublicKey         string `json:"publicKey"`
}
