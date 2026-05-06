package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/vw"
)

var pipelinesCmds = cmd.CommandGroup("",
	cmd.CommandGroup("gitlab",
		&GitlabPipelinesCmd,
		&GitlabPipelineCmd,
		&GitlabPipelineJobsCmd,
		&JobLogCmd,
		&GitlabRegistryCmd,
		&GitlabMergeRequestsCmd,
		&GitLabMergeRequestCmd,
		&MergeRequestCheckCmd,
		&GitlabAutoHealMergeCmd,
	),
	cmd.CommandGroup("fortify",
		&FortifyFalsePositiveCmd,
	),
	cmd.CommandGroup("info",
		&InfoProjectIdsCmd,
		&InfoProjectGroupsCmd,
	),
	cmd.CommandGroup("api",
		&SimplApiTokenizeCmd,
		&SimplApiEchoCmd,
	),
)

var views = map[string]vw.View{
	"table": vw.NewTableWriter(os.Stdout),
	"json":  vw.NewJsonView(os.Stdout),
}

func main() {
	fields := ""
	viewf := ""
	fl := flag.NewFlagSet("", flag.ExitOnError)
	fl.StringVar(&fields, "f", "", "")
	fl.StringVar(&viewf, "o", "table", "")
	fl.Parse(os.Args[1:])

	view, ok := views[viewf]
	if !ok {
		fmt.Fprintf(os.Stderr, "invalid -o parameter %q\n", viewf)
		os.Exit(1)
	}

	res, err := pipelinesCmds.CRun(fl.Args())
	if err != nil && !errors.Is(err, cmd.MissingCommand) {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	v := reflect.ValueOf(res)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() == reflect.Int {
		os.Exit(v.Interface().(int))
	}

	view.Render(&vw.RenderOpt{
		Fields: getFields(fields),
	}, res)

	os.Exit(0)
}
