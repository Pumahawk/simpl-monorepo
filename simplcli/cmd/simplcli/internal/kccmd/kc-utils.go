package kccmd

import "github.com/Pumahawk/simpl-monorepo/internal/kc"

type aCFT struct {
	User    string
	Pass    string
	BaseUrl string
	Realm   string
}

var acf = &aCFT{}

func (s *aCFT) NewClient() *kc.Client {
	return &kc.Client{
		BaseUrl: s.BaseUrl,
		AuthFunc: func() (*kc.AuthInfo, error) {
			return &kc.AuthInfo{
				Username:  s.User,
				Passaword: s.Pass,
				Realm:     s.Realm,
			}, nil
		},
	}
}
