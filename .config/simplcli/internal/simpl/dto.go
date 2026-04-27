package simpl

// Check health DTO
type healtResponseDto struct {
	Status string
}

type GenerateKeypairRequestDto struct {
	Name string `json:"name"`
}

type GenerateKeypairResponseDto struct {
	Name string `json:"name"`
}
