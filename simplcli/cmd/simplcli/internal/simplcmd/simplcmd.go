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
		&KeypairActiveCmd,
		&NewKeyPair,
	},
	FlagFunc: func(fs *flag.FlagSet) {
		fs.StringVar(&sacf.User, "user", utils.EnvOrDef("SPUSER", ""), "") // Default user is defined to the API level
		fs.StringVar(&sacf.Pass, "pass", utils.EnvOrDef("SPPASWORD", "password"), "")
		fs.StringVar(&sacf.BaseUrl, "baseurl", utils.EnvOrDef("SPBASEURL", "http://localhost:8100"), "")
		fs.StringVar(&sacf.Realm, "realm", utils.EnvOrDef("SPREALM", "authority"), "")
		fs.BoolVar(&sacf.KubeProxy, "kube", false, "")

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
