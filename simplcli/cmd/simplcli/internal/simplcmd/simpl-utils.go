package simplcmd

import (
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/simpl"
)

type simplACFT struct {
	User  string
	Pass  string
	Realm string
}

var SimplEndpoint = &simpl.SingleAccessEndpoints{}

func (s *simplACFT) NewClient(defUsr string) *simpl.Client {
	user := s.User
	if user == "" {
		user = defUsr
	}
	return &simpl.Client{
		Endpoints: SimplEndpoint,
		AuthFunc: func() (*simpl.AuthInfo, error) {
			return &simpl.AuthInfo{
				Username:  user,
				Passaword: s.Pass,
				Realm:     s.Realm,
			}, nil
		},
	}
}
