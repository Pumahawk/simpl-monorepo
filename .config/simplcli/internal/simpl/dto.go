package simpl

// Check health DTO
type healtResponseDto struct {
	Status string `json:"status"`
}

type GenerateKeypairRequestDto struct {
	Name string `json:"name"`
}

type GenerateKeypairResponseDto struct {
	Csr               string `json:"csr"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	Active            bool   `json:"active"`
	PublicKey         string `json:"publicKey"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type KeypairsSearch struct {
	Name string `search:"name"`
}

type KeypairsResponseDto struct {
	Items []KeypairsItemResponseDto `json:"items"`
}

type KeypairsItemResponseDto struct {
	Name              string `json:"name"`
	CreationTimestamp string `json:"creationTimestamp"`
	Active            bool   `json:"active"`
	Id                string `json:"id"`
	PublicKey         string `json:"publicKey"`
	Csr               string `json:"csr"`
}

type ParticipantsSearch struct {
	Organization string `search:"organization"`
}

type ParticipantsResponseDto struct {
	Items []ParticipantsItemResponseDto `json:"items"`
}

type ParticipantsItemResponseDto struct {
	Id                string `json:"id"`
	Organization      string `json:"organization"`
	CreationTimestamp string `json:"creationTimestamp"`
	UpdateTimestamp   string `json:"updateTimestamp"`
	TierOnePublicKey  string `json:"tierOnePublicKey"`
}
