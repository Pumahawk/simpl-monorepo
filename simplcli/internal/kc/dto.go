package kc

type RealmsResponseDto struct {
	AccountServiceService    string `json:"account-service"`
	PublicKey                string `json:"public_key"`
	Realm                    string `json:"realm"`
	TokenServiceService      string `json:"token-service"`
	TokensNotBeforeNotBefore int    `json:"tokens-not-before"`
}
