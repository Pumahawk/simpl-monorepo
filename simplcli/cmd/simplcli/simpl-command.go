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

		fl := flag.NewFlagSet("", flag.ExitOnError)
		sacf := simplApiCommonFlags(fl)
		fl.Parse(args)

		cl := sacf.NewClient()

		if err := simplApiCommonFlagsValidator(sacf); err != nil {
			return nil, err
		}

		return cl.Echo()
	},
}
