package simplcmd

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
)

var sacf = &simplACFT{}

var Cmd = &cmd.CommandGroup{
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
