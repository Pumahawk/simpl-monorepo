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
		sacf := &simplACFT{}
		structFlag(fs, sacf)
	},
	FlagValFunc: func() error {
		if sacf.User == "" {
			return fmt.Errorf("missing user flag")
		}

		if sacf.Pass == "" {
			return fmt.Errorf("missing pass flag")
		}

		if sacf.Server == "" {
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

		cl := sacf.NewClient()

		token, err := cl.Tokenize()
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
		cl := sacf.NewClient()
		return cl.Echo()
	},
}
