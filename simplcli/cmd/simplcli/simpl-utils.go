package main

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/internal/simpl"
)

type simplACFT struct {
	User   string
	Pass   string
	Server string
	Realm  string
}

func (s *simplACFT) NewClient() *simpl.Client {
	return &simpl.Client{
		BaseUrl: s.Server,
		AuthFunc: func() (*simpl.AuthInfo, error) {
			return &simpl.AuthInfo{
				Username:  s.User,
				Passaword: s.Pass,
				Realm:     s.Realm,
			}, nil
		},
	}

}

func simplApiCommonFlags(fs *flag.FlagSet) *simplACFT {
	sacf := &simplACFT{}
	structFlag(fs, sacf)
	return sacf
}

func simplApiCommonFlagsValidator(simplApiCommonFlags *simplACFT) error {
	if simplApiCommonFlags.User == "" {
		return fmt.Errorf("missing user flag")
	}

	if simplApiCommonFlags.Pass == "" {
		return fmt.Errorf("missing pass flag")
	}

	if simplApiCommonFlags.Server == "" {
		return fmt.Errorf("missing server flag")
	}

	if simplApiCommonFlags.Realm == "" {
		return fmt.Errorf("missing realm flag")
	}
	return nil
}
