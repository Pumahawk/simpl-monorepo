package main

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var GitlabPipelinesCmd = Command[*gitlab.PipelinesResponseDto]{
	Name: "pip",
	Run: func(c *Command[*gitlab.PipelinesResponseDto], args []string) (*gitlab.PipelinesResponseDto, error) {
		search := &gitlab.SearchPipeline{}

		fl := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fl, search)
		fl.Parse(args)
		projectId := fl.Arg(0)

		if projectId == "" {
			return nil, fmt.Errorf("missing project id")
		}

		res, err := gitlabClient.Pipelines(prIds.Get(projectId), search)
		if err != nil {
			return nil, fmt.Errorf("error get project %s: %s", projectId, err)
		}

		return res, nil
	},
}

var GitlabPipelineCmd = Command[*gitlab.PipelineResponseDto]{
	Name: "pipd",
	Run: func(c *Command[*gitlab.PipelineResponseDto], args []string) (*gitlab.PipelineResponseDto, error) {
		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.Parse(args)
		projectId := fl.Arg(0)
		pipelineId := fl.Arg(1)

		if projectId == "" {
			return nil, fmt.Errorf("missing project id")
		}

		if pipelineId == "" {
			return nil, fmt.Errorf("missing pipeline id")
		}

		r, err := gitlabClient.Pipeline(prIds.Get(projectId), pipelineId)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve pipeline %q project %q: %w", pipelineId, projectId, err)
		}

		return r, nil
	},
}

var GitlabPipelineJobsCmd = Command[any]{
	Name: "pipj",
	Run: func(c *Command[any], args []string) (any, error) {
		search := &gitlab.SearchPipelineJob{}

		fl := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fl, search)
		fl.Parse(args)
		projectId := fl.Arg(0)
		pipelineId := fl.Arg(1)

		if projectId == "" {
			return nil, fmt.Errorf("missing project id")
		}

		if pipelineId == "" {
			return nil, fmt.Errorf("missing pipeline id")
		}

		r, err := gitlabClient.PipelineJobs(prIds.Get(projectId), pipelineId, search)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve pipeline jobs %q project %q: %w", pipelineId, projectId, err)
		}

		return r, nil
	},
}
