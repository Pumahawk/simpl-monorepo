package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var GitlabPipelinesCmd = Command{
	Name: "pip",
	Run: func(c *Command, args []string) int {
		search := &gitlab.SearchPipeline{}
		fields := ""

		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.StringVar(&fields, "f", "", "")
		fl.StringVar(&search.Ref, "ref", "", "")
		fl.StringVar(&search.Status, "status", "", "")
		fl.StringVar(&search.Page, "page", "", "")
		fl.StringVar(&search.PerPage, "size", "", "")
		fl.Parse(args)
		projectId := fl.Arg(0)

		if projectId == "" {
			fmt.Fprintln(os.Stderr, "missing project id")
			return 1
		}

		res, err := gitlabClient.Pipelines(prIds.Get(projectId), search)
		if err != nil {
			log.Fatalf("error get project %s: %s", projectId, err)
		}

		vw := StdTableWriter()
		vw.RenderList(&RenderOpt{
			Fields: getFields(fields),
		}, res)

		return 0
	},
}

var GitlabPipelineCmd = Command{
	Name: "pipd",
	Run: func(c *Command, args []string) int {
		fields := ""

		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.StringVar(&fields, "f", "", "")
		fl.Parse(args)
		projectId := fl.Arg(0)
		pipelineId := fl.Arg(1)

		if projectId == "" {
			fmt.Fprintln(os.Stderr, "missing project id")
			return 1
		}

		if pipelineId == "" {
			fmt.Fprintln(os.Stderr, "missing pipeline id")
			return 1
		}

		r, err := gitlabClient.Pipeline(prIds.Get(projectId), pipelineId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to retrieve pipeline %q project %q", pipelineId, projectId)
			return 1
		}

		wm := StdTableWriter()
		wm.RenderValue(&RenderOpt{
			Fields: getFields(fields),
		}, r)
		return 0
	},
}
