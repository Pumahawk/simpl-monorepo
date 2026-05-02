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
	Name: "pipeline:search",
	Run: func(c *cmd.Command[[]gitlab.PipelineResponseItemDto], args []string) ([]gitlab.PipelineResponseItemDto, error) {
		search := &gitlab.SearchPipeline{}

		fl := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fl, search)
		fl.Parse(args)
		projectIds := fl.Args()

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
		ress := make([]*resTy, 0, len(projectIds))

		// Retrieve all pipelines async
		for _, projectId := range projectIds {
			res := &resTy{}
			res.name = projectId
			ress = append(ress, res)
			wg.Add(1)
			go func(projectId string) {
				defer wg.Done()
				res.rs, res.err = gitlabClient.Pipelines(prIds.Get(projectId), search)
			}(projectId)
		}
		wg.Wait()

		// Prepare items size slice
		size := 0
		for _, r := range ress {
			if r.rs != nil {
				size += len(r.rs.Items)
			}
		}
		items := make([]gitlab.PipelineResponseItemDto, 0, size)

		// Collect all response items
		// Rewrite pipeline name with project name.
		for _, r := range ress {
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
