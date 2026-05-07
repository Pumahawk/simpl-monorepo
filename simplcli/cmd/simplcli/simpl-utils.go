package main

import (
	"github.com/Pumahawk/simpl-monorepo/internal/simpl"
)

type simplACFT struct {
	User    string
	Pass    string
	BaseUrl string
	Realm   string
}

func (s *simplACFT) NewClient(defUsr string) *simpl.Client {
	user := s.User
	if user == "" {
		user = defUsr
	}
	return &simpl.Client{
		BaseUrl: s.BaseUrl,
		AuthFunc: func() (*simpl.AuthInfo, error) {
			return &simpl.AuthInfo{
				Username:  user,
				Passaword: s.Pass,
				Realm:     s.Realm,
			}, nil
		},
	}

}
