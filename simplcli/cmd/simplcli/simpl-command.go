package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/simpl"
)

var sacf = &simplACFT{}

var simplCmdG = &cmd.CommandGroup{
	Name: "api",
	Commands: []cmd.CommandW{
		&SimplApiTokenizeCmd,
		&SimplApiEchoCmd,
	},
	FlagFunc: func(fs *flag.FlagSet) {
		fs.StringVar(&sacf.User, "user", "", "") // Default user is defined to the API level
		fs.StringVar(&sacf.Pass, "pass", "password", "")
		fs.StringVar(&sacf.BaseUrl, "baseurl", "http://localhost:8100", "")
		fs.StringVar(&sacf.Realm, "realm", "authority", "")

	},
	FlagValFunc: func() error {

		if sacf.Pass == "" {
			return fmt.Errorf("missing pass flag")
		}

		if sacf.BaseUrl == "" {
			return fmt.Errorf("missing server flag")
		}

		if sacf.Realm == "" {
			return fmt.Errorf("missing realm flag")
		}
		return nil
	},
}

var SimplApiTokenizeCmd = cmd.Command[int]{
	Name: "tokenize",
	Run: func(c *cmd.Command[int], args []string) (int, error) {

		cl := sacf.NewClient("m.m")

		token, err := cl.Tokenize()
		if err != nil {
			return 1, fmt.Errorf("unable to tokenize server=%q, user=%q: %w", sacf.BaseUrl, sacf.User, err)
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
		cl := sacf.NewClient("m.m")
		return cl.Echo()
	},
}
