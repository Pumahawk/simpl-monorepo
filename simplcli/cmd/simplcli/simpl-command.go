package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/simpl"
)

var SimplApiTokenizeCmd = cmd.Command[int]{
	Name: "tokenize",
	Run: func(c *cmd.Command[int], args []string) (int, error) {

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
			return 1, fmt.Errorf("missing user flag")
		}

		if pass == "" {
			return 1, fmt.Errorf("missing pass flag")
		}

		if server == "" {
			return 1, fmt.Errorf("missing server flag")
		}

		if realm == "" {
			return 1, fmt.Errorf("missing realm flag")
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
			return 1, fmt.Errorf("unable to tokenize server=%q, user=%q: %w", server, user, err)
		}

		_, err = os.Stdout.Write([]byte(token.AccessToken))
		if err != nil {
			return 1, fmt.Errorf("unable to write to stdout")
		}

		return 0, nil
	},
}
