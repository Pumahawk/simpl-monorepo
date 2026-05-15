package main

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/gitlab"
)

type GitlabPipelinesArgs struct {
	My bool
}

var GitlabPipelinesCmd = SearchMultiProjectAsyncCmd[gitlab.SearchPipeline, GitlabPipelinesArgs, gitlab.PipelineResponseItemDto]{
	Name: "pipelines:search",
	ApiFunc: func(projectId string, flags GitlabPipelinesArgs, search *gitlab.SearchPipeline) ([]gitlab.PipelineResponseItemDto, error) {
		if flags.My {
			user, err := gitlabClient.CurrentUser()
			if err != nil {
				return nil, fmt.Errorf("retrieve current user: %w", err)
			}
			search.Username = user.Username
		}
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

type GitlabPipelineModel struct {
	Details    *gitlab.PipelineResponseDto
	Attributes []gitlab.PipelineAttributesItemResponseDto
}

var GitlabPipelineCmd = cmd.Command[*GitlabPipelineModel]{
	Name: "pipelines:details",
	Run: func(c *cmd.Command[*GitlabPipelineModel], args []string) (*GitlabPipelineModel, error) {
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

		type getPipelineResult struct {
			r *gitlab.PipelineResponseDto
			error
		}
		getPipelineResultC := make(chan getPipelineResult, 1)
		go func() {
			r, err := gitlabClient.Pipeline(prIds.Get(projectId), pipelineId)
			if err != nil {
				getPipelineResultC <- getPipelineResult{nil, fmt.Errorf("unable to retrieve pipeline %q project %q: %w", pipelineId, projectId, err)}
			} else {
				getPipelineResultC <- getPipelineResult{r, nil}
			}
		}()

		type getPipelineAttributesResult struct {
			r *gitlab.PipelineAttributesResponseDto
			error
		}
		getPipelineAttributesResultC := make(chan getPipelineAttributesResult, 1)
		go func() {
			r, err := gitlabClient.PipelineAttributes(prIds.Get(projectId), pipelineId)
			if err != nil {
				getPipelineAttributesResultC <- getPipelineAttributesResult{nil, fmt.Errorf("unable to retrieve pipeline %q attributes project %q: %w", pipelineId, projectId, err)}
			} else {
				getPipelineAttributesResultC <- getPipelineAttributesResult{r, nil}
			}
		}()

		details := <-getPipelineResultC
		attributes := <-getPipelineAttributesResultC
		if err := details.error; err != nil {
			return nil, err
		}
		if err := attributes.error; err != nil {
			return nil, err
		}

		return &GitlabPipelineModel{
			Details:    details.r,
			Attributes: attributes.r.Items,
		}, nil
	},
}

var GitlabPipelineJobsCmd = cmd.Command[[]gitlab.PipelineJobsResponseItemDto]{
	Name: "pipelines:jobs",
	Run: func(c *cmd.Command[[]gitlab.PipelineJobsResponseItemDto], args []string) ([]gitlab.PipelineJobsResponseItemDto, error) {
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

		return r.Items, nil
	},
}

var GitlabMergeRequestsCmd = SearchMultiProjectAsyncCmd[gitlab.SearchMergeRequest, struct{}, gitlab.MergeRequestsResponseItemDto]{
	Name: "merge:search",
	ApiFunc: func(projectId string, _ struct{}, search *gitlab.SearchMergeRequest) ([]gitlab.MergeRequestsResponseItemDto, error) {
		r, err := gitlabClient.MergeRequests(projectId, search)
		if err != nil {
			return nil, err
		}
		return r.Items, nil
	},
	SortFunc: func(mrrid []gitlab.MergeRequestsResponseItemDto, i, j int) bool {
		return mrrid[i].UpdatedAt < mrrid[j].UpdatedAt
	},
}

var GitLabMergeRequestCmd = cmd.Command[*gitlab.MergeRequestResponseDto]{
	Name: "merge:details",
	Run: func(c *cmd.Command[*gitlab.MergeRequestResponseDto], args []string) (*gitlab.MergeRequestResponseDto, error) {
		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.Parse(args)
		projectId := fl.Arg(0)
		mergeRequestId := fl.Arg(1)

		if projectId == "" {
			return nil, fmt.Errorf("missing project id")
		}

		if mergeRequestId == "" {
			return nil, fmt.Errorf("missing merge request id")
		}

		r, err := gitlabClient.MergeRequest(prIds.Get(projectId), mergeRequestId)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve merger request %q project %q: %w", mergeRequestId, projectId, err)
		}

		return r, nil
	},
}
