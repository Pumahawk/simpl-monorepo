package kc

type RealmResponseDto struct {
	AccountServiceService    string `json:"account-service"`
	PublicKey                string `json:"public_key"`
	Realm                    string `json:"realm"`
	TokenServiceService      string `json:"token-service"`
	TokensNotBeforeNotBefore int    `json:"tokens-not-before"`
}

type RealmsItemResponseDto struct {
	Id    string `json:"Id"`
	Realm string `json:"Realm"`
}
