package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var GitlabPipelinesCmd = cmd.Command[[]gitlab.PipelineResponseItemDto]{
	Name: "pipelines:search",
	Run: func(c *cmd.Command[[]gitlab.PipelineResponseItemDto], args []string) ([]gitlab.PipelineResponseItemDto, error) {
		search := &gitlab.SearchPipeline{}

		fl := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fl, search)
		fl.Parse(args)
		projectIds := prIdsDemux.demux(fl.Args())

		if len(projectIds) == 0 {
			return nil, fmt.Errorf("missing project ids")
		}

		// Prepare response object container
		type resTy struct {
			name string
			rs   *gitlab.PipelinesResponseDto
			err  error
		}
		wg := sync.WaitGroup{}
		ress := make(chan resTy)

		// Retrieve all pipelines async
		for _, projectId := range projectIds {
			name := projectId
			wg.Go(func() {
				rs, err := gitlabClient.Pipelines(prIds.Get(projectId), search)
				ress <- resTy{
					name: name,
					rs:   rs,
					err:  err,
				}
			})
		}
		go func() {
			defer close(ress)
			wg.Wait()
		}()

		// Collect all response items
		// Rewrite pipeline name with project name.
		items := make([]gitlab.PipelineResponseItemDto, 0, len(projectIds)*20)
		for r := range ress {
			if r.err != nil {
				fmt.Fprintf(os.Stderr, "projectId %s: %s\n", r.name, r.err)
			}
			if r.rs != nil {
				for _, item := range r.rs.Items {
					item.Name = r.name
					items = append(items, item)
				}
			}
		}

		// Sort items by updated_at
		sort.Slice(items, func(i, j int) bool {
			return items[i].UpdatedAt > items[j].UpdatedAt
		})

		return items, nil
	},
}

var GitlabPipelineCmd = cmd.Command[*gitlab.PipelineResponseDto]{
	Name: "pipelines:details",
	Run: func(c *cmd.Command[*gitlab.PipelineResponseDto], args []string) (*gitlab.PipelineResponseDto, error) {
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

var GitlabPipelineJobsCmd = cmd.Command[any]{
	Name: "pipelines:jobs",
	Run: func(c *cmd.Command[any], args []string) (any, error) {
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
