package kccmd

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/kc"
)

var Cmd = &cmd.CommandGroup{
	Name: "keycloak",
	Commands: []cmd.CommandW{
		&RealmGetCmd,
		&RealmLsCmd,
	},
	FlagFunc: func(fs *flag.FlagSet) {
		fs.StringVar(&acf.User, "user", envOrDef("KCUSER", "admin"), "")
		fs.StringVar(&acf.Pass, "pass", envOrDef("KCPASSWORD", "admin"), "")
		fs.StringVar(&acf.BaseUrl, "baseurl", envOrDef("KCBASEURL", "http://localhost:8100/auth"), "")
		fs.StringVar(&acf.Realm, "realm", envOrDef("KCREALM", "master"), "")
	},
}

var RealmLsCmd = cmd.Command[[]kc.RealmsItemResponseDto]{
	Name: "realms:list",
	Run: func(c *cmd.Command[[]kc.RealmsItemResponseDto], args []string) ([]kc.RealmsItemResponseDto, error) {
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.Parse(args)

		rs := acf.NewClient()
		res, err := rs.Realms()
		if err != nil {
			return nil, fmt.Errorf("list realms %w", err)
		}
		return res, nil
	},
}

var RealmGetCmd = cmd.Command[*kc.RealmResponseDto]{
	Name: "realms:details",
	Run: func(c *cmd.Command[*kc.RealmResponseDto], args []string) (*kc.RealmResponseDto, error) {
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.Parse(args)

		realm := fs.Arg(0)

		if realm == "" {
			return nil, fmt.Errorf("missing realm")
		}

		rs := acf.NewClient()
		res, err := rs.Realm(realm)
		if err != nil {
			return nil, fmt.Errorf("get realm %q: %w", realm, err)
		}
		return res, nil
	},
}
