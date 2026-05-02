package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var MergeRequestCheckCmd = cmd.Command[MergeRequestCheckModel]{
	Name: "pipelines:check",
	Run: func(c *cmd.Command[MergeRequestCheckModel], args []string) (MergeRequestCheckModel, error) {
		search := &gitlab.SearchPipeline{}

		fs := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fs, search)
		fs.Parse(args)
		projectIds := fs.Args()

		if len(projectIds) == 0 {
			return nil, fmt.Errorf("missing projectId")
		}

		// Retrieve all pipelines and jobs async
		type plsj struct {
			projectName string
			p           *gitlab.PipelineResponseItemDto
			r           *gitlab.PipelineJobsResponseDto
			err         error
		}
		plsjl := make(chan plsj)
		wg := sync.WaitGroup{}
		for _, projectId := range projectIds {
			projectName := projectId
			projectId := prIds.Get(projectId)
			// Get pipeline informations
			wg.Go(func() {
				pls, err := gitlabClient.Pipelines(projectId, search)
				if err != nil {
					plsjl <- plsj{
						projectName: projectName,
						p:           nil,
						r:           nil,
						err:         fmt.Errorf("get pipelinese projectId=%q: %w", projectId, err),
					}
					return
				}

				for _, _pl := range pls.Items {
					pl := &_pl
					// Get job informations
					wg.Go(func() {
						r, err := gitlabClient.PipelineJobs(projectId, strconv.Itoa(pl.Id), &gitlab.SearchPipelineJob{})
						plsjl <- plsj{
							projectName: projectName,
							p:           pl,
							r:           r,
							err:         err,
						}
					})
				}
			})
		}

		// Wait all goroutines end and close channel
		go func() {
			defer close(plsjl)
			wg.Wait()
		}()

		// Map model
		model := make([]MRChPipeline, 0, 50)
		for res := range plsjl { // Channel closed by dedicated goroutine
			if res.p == nil || res.err != nil {
				if res.p == nil {
					fmt.Fprintf(os.Stderr, "unable to retrieve pipeline: %s\n", res.err)
				} else {
					fmt.Fprintf(os.Stderr, "unable to retrieve pipeline jobs %q\n: %s", res.p.Id, res.err)
				}
			} else {
				// Count all success and failed jobs
				// Collect failed job names
				jobst := len(res.r.Items)
				jobsok := 0
				jobses := make([]string, 0, len(res.r.Items))
				for _, j := range res.r.Items {
					switch j.Status {
					case "success":
						jobsok++
					case "failed":
						jobses = append(jobses, fmt.Sprintf("[%s:%d]", j.Name, j.Id))
					}
				}
				jobs := fmt.Sprintf("%d/%d", jobsok, jobst)
				model = append(model, MRChPipeline{
					Id:         res.p.Id,
					Project:    res.projectName,
					Status:     res.p.Status,
					Ref:        res.p.Ref,
					Source:     res.p.Source,
					UpdatedAt:  res.p.UpdatedAt,
					Jobs:       jobs,
					JobsErrors: strings.Join(jobses, ","),
				})
			}
		}

		sort.Slice(model, func(i, j int) bool {
			return model[i].UpdatedAt < model[j].UpdatedAt
		})

		return MergeRequestCheckModel(model), nil
	},
}

type MergeRequestCheckModel []MRChPipeline

type MRChPipeline struct {
	Id         int // pipelineId
	Project    string
	Status     string
	Ref        string
	Source     string
	UpdatedAt  string
	Jobs       string // jobs count ex: 10/12
	JobsErrors string // jobs error names separated by comma, ex: [job1:job1Id],[job3:job3Id],[job6:job6Id]
}
