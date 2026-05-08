package kccmd

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/kc"
)

var Cmd = &cmd.CommandGroup{
	Name: "keycloak",
	Commands: []cmd.CommandW{
		&RealmGetCmd,
	},
	FlagFunc: func(fs *flag.FlagSet) {
		fs.StringVar(&acf.User, "user", "admin", "")
		fs.StringVar(&acf.Pass, "pass", "admin", "")
		fs.StringVar(&acf.BaseUrl, "baseurl", "http://localhost:8100/auth", "")
		fs.StringVar(&acf.Realm, "realm", "master", "")
	},
}

var RealmGetCmd = cmd.Command[*kc.RealmsResponseDto]{
	Name: "realms:details",
	Run: func(c *cmd.Command[*kc.RealmsResponseDto], args []string) (*kc.RealmsResponseDto, error) {
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.Parse(args)

		realm := fs.Arg(0)

		if realm == "" {
			return nil, fmt.Errorf("missing realm")
		}

		rs := acf.NewClient()
		res, err := rs.Realms(realm)
		if err != nil {
			return nil, fmt.Errorf("get realm %q: %w", realm, err)
		}
		return res, nil
	},
}
