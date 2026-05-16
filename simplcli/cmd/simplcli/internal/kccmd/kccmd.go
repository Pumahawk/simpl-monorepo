package kccmd

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/kc"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/utils"
)

var Cmd = &cmd.CommandGroup{
	Name: "keycloak",
	Commands: []cmd.CommandW{
		&RealmGetCmd,
		&RealmLsCmd,
		&RealmExportCmd,
		&RealmImportCmd,
		&RealmDeleteCmd,
	},
	FlagFunc: func(fs *flag.FlagSet) {
		fs.StringVar(&acf.User, "user", utils.EnvOrDef("KCUSER", "admin"), "")
		fs.StringVar(&acf.Pass, "pass", utils.EnvOrDef("KCPASSWORD", "admin"), "")
		fs.StringVar(&acf.BaseUrl, "baseurl", utils.EnvOrDef("KCBASEURL", "http://localhost:8100/auth"), "")
		fs.StringVar(&acf.ClientId, "clientid", utils.EnvOrDef("KCCLIENTID", "admin-cli"), "")
		fs.StringVar(&acf.Realm, "realm", utils.EnvOrDef("KCREALM", "master"), "")
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

var RealmDeleteCmd = cmd.Command[int]{
	Name: "realms:delete",
	Run: func(c *cmd.Command[int], args []string) (int, error) {
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.Parse(args)

		realm := fs.Arg(0)

		if realm == "" {
			return 1, fmt.Errorf("missing realm")
		}

		rs := acf.NewClient()
		err := rs.RealmDelete(realm)
		if err != nil {
			return 1, fmt.Errorf("delete realm %q: %w", realm, err)
		}
		return 0, nil
	},
}

var RealmExportCmd = cmd.Command[int]{
	Name: "realms:export",
	Run: func(c *cmd.Command[int], args []string) (int, error) {
		opt := &kc.RealmExportOpt{}
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.BoolVar(&opt.IncludeClients, "include-clients", false, "")
		fs.BoolVar(&opt.ExportGroupsAndRoles, "export-groups-and-roles", false, "")
		fs.Parse(args)

		realm := fs.Arg(0)

		if realm == "" {
			return 1, fmt.Errorf("missing realm")
		}

		rs := acf.NewClient()
		res, err := rs.RealmExport(realm, opt)
		if err != nil {
			return 1, fmt.Errorf("export realm %q: %w", realm, err)
		}

		io.Copy(os.Stdout, bytes.NewBuffer(res))
		return 0, nil
	},
}

var RealmImportCmd = cmd.Command[int]{
	Name: "realms:import",
	Run: func(c *cmd.Command[int], args []string) (int, error) {
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.Parse(args)

		realm := fs.Arg(0)

		doc := &bytes.Buffer{}
		_, err := io.Copy(doc, os.Stdin)
		if err != nil {
			return 1, err
		}

		rs := acf.NewClient()
		err = rs.RealmImport(doc.Bytes())
		if err != nil {
			return 1, fmt.Errorf("import realm %q: %w", realm, err)
		}

		return 0, nil
	},
}
