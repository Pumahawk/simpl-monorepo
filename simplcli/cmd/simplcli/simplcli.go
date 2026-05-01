package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
)

var pipelinesCmds = CommandGroup("",
	CommandGroup("gl",
		&GitlabPipelinesCmd,
		&GitlabPipelineCmd,
		&GitlabPipelineJobsCmd,
	),
	CommandGroup("info",
		&InfoProjectIdsCmd,
	),
)

func main() {
	fields := ""
	fl := flag.NewFlagSet("", flag.ExitOnError)
	fl.StringVar(&fields, "f", "", "")
	fl.Parse(os.Args[1:])

	res, err := pipelinesCmds.CRun(fl.Args())
	if err != nil && !errors.Is(err, MissingCommand) {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	v := reflect.ValueOf(res)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() == reflect.Int {
		os.Exit(*(v.Interface().(*int)))
	}

	vw := StdTableWriter()
	vw.Render(&RenderOpt{
		Fields: getFields(fields),
	}, res)

	os.Exit(0)
}
