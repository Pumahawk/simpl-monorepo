package main

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var GitlabPipelinesCmd = SearchMultiProjectAsyncCmd[gitlab.SearchPipeline, struct{}, gitlab.PipelineResponseItemDto]{
	Name: "pipelines:search",
	ApiFunc: func(projectId string, flags struct{}, search *gitlab.SearchPipeline) ([]gitlab.PipelineResponseItemDto, error) {
		r, err := gitlabClient.Pipelines(projectId, search)
		if err != nil {
			return nil, err
		}
		return r.Items, nil
	},
	SortFunc: func(prid []gitlab.PipelineResponseItemDto, i, j int) bool {
		return prid[i].UpdatedAt < prid[j].UpdatedAt
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

var GitlabMergeRequestsCmd = SearchMultiProjectAsyncCmd[gitlab.SearchMergeRequest, struct{}, gitlab.MergeRequestResponseItemDto]{
	Name: "merge:search",
	ApiFunc: func(projectId string, _ struct{}, search *gitlab.SearchMergeRequest) ([]gitlab.MergeRequestResponseItemDto, error) {
		r, err := gitlabClient.MergeRequests(projectId, search)
		if err != nil {
			return nil, err
		}
		return r.Items, nil
	},
	SortFunc: func(mrrid []gitlab.MergeRequestResponseItemDto, i, j int) bool {
		return mrrid[i].UpdatedAt < mrrid[j].UpdatedAt
	},
}
