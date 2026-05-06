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

		fl := flag.NewFlagSet("", flag.ExitOnError)
		sacf := simplApiCommonFlags(fl)
		fl.Parse(args)

		if err := simplApiCommonFlagsValidator(sacf); err != nil {
			return 1, err
		}

		client := &simpl.Client{
			BaseUrl: sacf.Server,
			AuthFunc: func() (*simpl.AuthInfo, error) {
				return &simpl.AuthInfo{
					Username:  sacf.User,
					Passaword: sacf.Pass,
					Realm:     sacf.Realm,
				}, nil
			},
		}

		token, err := client.Tokenize()
		if err != nil {
			return 1, fmt.Errorf("unable to tokenize server=%q, user=%q: %w", sacf.Server, sacf.User, err)
		}

		_, err = os.Stdout.Write([]byte(token.AccessToken))
		if err != nil {
			return 1, fmt.Errorf("unable to write to stdout")
		}

		return 0, nil
	},
}

var SimplApiEchoCmd = cmd.Command[*simpl.EchoResponseDto]{
	Name: "echo",
	Run: func(c *cmd.Command[*simpl.EchoResponseDto], args []string) (*simpl.EchoResponseDto, error) {

		fl := flag.NewFlagSet("", flag.ExitOnError)
		sacf := simplApiCommonFlags(fl)
		fl.Parse(args)

		if err := simplApiCommonFlagsValidator(sacf); err != nil {
			return nil, err
		}

		client := &simpl.Client{
			BaseUrl: sacf.Server,
			AuthFunc: func() (*simpl.AuthInfo, error) {
				return &simpl.AuthInfo{
					Username:  sacf.User,
					Passaword: sacf.Pass,
					Realm:     sacf.Realm,
				}, nil
			},
		}
		return client.Echo()
	},
}

type simplACFT struct {
	User   string
	Pass   string
	Server string
	Realm  string
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
