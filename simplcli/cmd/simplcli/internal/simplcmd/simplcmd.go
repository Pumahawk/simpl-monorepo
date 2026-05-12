package simplcmd

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/utils"
)

var sacf = &simplACFT{}

var Cmd = &cmd.CommandGroup{
	Name: "api",
	Commands: []cmd.CommandW{
		&SimplApiTokenizeCmd,
		&SimplApiEchoCmd,
	},
	FlagFunc: func(fs *flag.FlagSet) {
		fs.StringVar(&sacf.User, "user", utils.EnvOrDef("SPUSER", ""), "") // Default user is defined to the API level
		fs.StringVar(&sacf.Pass, "pass", utils.EnvOrDef("SPPASWORD", "password"), "")
		fs.StringVar(&SimplEndpoint.BaseUrl, "baseurl", utils.EnvOrDef("SPBASEURL", "http://localhost:8100"), "")
		fs.StringVar(&sacf.Realm, "realm", utils.EnvOrDef("SPREALM", "authority"), "")

	},
	FlagValFunc: func() error {

		if sacf.Pass == "" {
			return fmt.Errorf("missing pass flag")
		}

		if SimplEndpoint.BaseUrl == "" {
			return fmt.Errorf("missing server flag")
		}

		if sacf.Realm == "" {
			return fmt.Errorf("missing realm flag")
		}
		return nil
	},
}
