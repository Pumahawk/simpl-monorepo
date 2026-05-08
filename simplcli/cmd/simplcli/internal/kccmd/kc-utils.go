package kccmd

import (
	"os"

	"github.com/Pumahawk/simpl-monorepo/internal/kc"
)

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

func envOrDef(key, def string) string {
	if e, ok := os.LookupEnv(key); ok {
		return e
	} else {
		return def
	}
}
