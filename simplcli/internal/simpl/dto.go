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

type GenerateKeyPairResponseDto struct {
	CreationTimestamp string `json:"creationTimestamp"`
	Name              string `json:"name"`
	Active            bool   `json:"active"`
	Id                string `json:"id"`
	PublicKey         string `json:"publicKey"`
}

type GenerateKeyPairRequestDto struct {
	Name string `json:"name"`
}

type DataspaceUpdateDto struct {
	Name         string `json:"name"`
	ContactEmail string `json:"contactEmail"`
}

type CsrRequest struct {
	CommonName         string `json:"commonName"`
	Country            string `json:"country"`
	Organization       string `json:"organization"`
	OrganizationalUnit string `json:"organizationalUnit"`
}

type ParticipantCreateRequestDto struct {
	Organization    string `json:"organization"`
	ParticipantType string `json:"participantType"`
	IsAuthority     bool   `json:"isAuthority"`
}
