package simplcmd

import (
	"fmt"
	"os"
	"sync"

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

var KeypairActiveCmd = cmd.Command[*simpl.KeyPairActiveDto]{
	Name: "kp-active",
	Run: func(c *cmd.Command[*simpl.KeyPairActiveDto], args []string) (*simpl.KeyPairActiveDto, error) {
		cl := sacf.NewClient("m.m")
		return cl.KeypairActive()
	},
}

var NewKeyPair = cmd.Command[[]simpl.GenerateKeyPairResponseDto]{
	Name: "kp-new",
	Run: func(c *cmd.Command[[]simpl.GenerateKeyPairResponseDto], args []string) ([]simpl.GenerateKeyPairResponseDto, error) {
		names := args
		if len(names) == 0 {
			return nil, fmt.Errorf("missing keypair name")
		}

		cl := sacf.NewClient("m.m")

		type resT struct {
			name string
			kp   *simpl.GenerateKeyPairResponseDto
			err  error
		}
		var keypairs <-chan resT

		{
			ks := make(chan resT, len(names))
			keypairs = ks
			go func() {
				wg := sync.WaitGroup{}
				defer close(ks)
				for _, name := range names {
					wg.Go(func() {
						r, err := cl.GenerateKeyPair(name)
						ks <- resT{name, r, err}
					})
				}
				wg.Wait()
			}()
		}

		res := make([]simpl.GenerateKeyPairResponseDto, 0, len(names))
		for r := range keypairs {
			if r.err != nil {
				fmt.Fprintf(os.Stderr, "error generate keypair %q: %s\n", r.name, r.err)
			} else {
				res = append(res, *r.kp)
			}
		}

		return res, nil
	},
}
