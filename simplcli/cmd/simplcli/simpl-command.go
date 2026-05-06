package main

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/simpl"
)

var SimplApiTokenizeCmd = cmd.Command[*simpl.TokenizeResponseDto]{
	Name: "tokenize",
	Run: func(c *cmd.Command[*simpl.TokenizeResponseDto], args []string) (*simpl.TokenizeResponseDto, error) {

		var user string
		var pass string
		var server string
		var realm string

		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.StringVar(&user, "user", "", "")
		fl.StringVar(&pass, "pass", "", "")
		fl.StringVar(&server, "server", "", "")
		fl.StringVar(&realm, "realm", "", "")
		fl.Parse(args)

		if user == "" {
			return nil, fmt.Errorf("missing user flag")
		}

		if pass == "" {
			return nil, fmt.Errorf("missing pass flag")
		}

		if server == "" {
			return nil, fmt.Errorf("missing server flag")
		}

		if realm == "" {
			return nil, fmt.Errorf("missing realm flag")
		}

		client := &simpl.Client{
			BaseUrl: server,
			AuthFunc: func() (*simpl.AuthInfo, error) {
				return &simpl.AuthInfo{
					Username:  user,
					Passaword: pass,
					Realm:     realm,
				}, nil
			},
		}

		token, err := client.Tokenize()
		if err != nil {
			return nil, fmt.Errorf("unable to tokenize server=%q, user=%q: %w", server, user, err)
		}

		return token, nil
	},
}
