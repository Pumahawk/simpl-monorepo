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
	Name: "check",
	Run: func(c *cmd.Command[MergeRequestCheckModel], args []string) (MergeRequestCheckModel, error) {
		search := &gitlab.SearchPipeline{}
		projectId := ""

		fs := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fs, search)
		fs.Parse(args)
		projectId = prIds.Get(fs.Arg(0))

		if projectId == "" {
			return nil, fmt.Errorf("missing projectId")
		}

		pls, err := gitlabClient.Pipelines(projectId, search)
		if err != nil {
			return nil, fmt.Errorf("get pipelinese projectId=%q: %w", projectId, err)
		}

		// Retrieve all pipeline jobs async
		type plsj struct {
			p   *gitlab.PipelineResponseItemDto
			r   *gitlab.PipelineJobsResponseDto
			err error
		}
		plsjl := make(chan plsj)
		wg := sync.WaitGroup{}
		for _, pl := range pls.Items {
			wg.Add(1)
			go func(pl *gitlab.PipelineResponseItemDto) {
				defer wg.Done()
				r, err := gitlabClient.PipelineJobs(projectId, strconv.Itoa(pl.Id), &gitlab.SearchPipelineJob{})
				plsjl <- plsj{
					p:   pl,
					r:   r,
					err: err,
				}
			}(&pl)
		}
		go func() {
			defer close(plsjl)
			wg.Wait()
		}()

		// Map model
		model := make([]MRChPipeline, 0, len(pls.Items))
		for res := range plsjl {
			if res.err != nil {
				fmt.Fprintf(os.Stderr, "unable to retrieve jobs pipeline %q\n", res.p.Id)
			} else if res.r != nil {
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
					Status:     res.p.Status,
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
	Status     string
	UpdatedAt  string
	Jobs       string // jobs count ex: 10/12
	JobsErrors string // jobs error names separated by comma, ex: [job1:job1Id],[job3:job3Id],[job6:job6Id]
}
