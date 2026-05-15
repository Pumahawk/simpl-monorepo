package simplcmd

import (
	"fmt"
	"os"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/simpl"
)

var SimplApiTokenizeCmd = cmd.Command[int]{
	Name: "tokenize",
	Run: func(c *cmd.Command[int], args []string) (int, error) {

		// TODO refactor
		cl := sacf.NewClient("m.m")

		token, err := cl.ApiToken()
		if err != nil {
			return 1, fmt.Errorf("unable to tokenize server=%q, user=%q: %w", cl.Client.Endpoints.Keycloak(), sacf.User, err)
		}

		_, err = os.Stdout.Write([]byte(token))
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
