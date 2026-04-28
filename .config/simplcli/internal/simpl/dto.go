package simpl

import "net/url"

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
	Id   string
	Name string
}

func (k *KeypairsSearch) Query(values url.Values) {
	if k.Name != "" {
		values.Add("name", k.Name)
	}
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
