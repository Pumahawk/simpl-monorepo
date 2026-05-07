package simplcmd

import (
	"fmt"
	"os"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/simpl"
)

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
