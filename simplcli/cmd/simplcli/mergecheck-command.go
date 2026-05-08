package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/gitlab"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/log"
)

var MergeRequestCheckCmd = cmd.Command[MergeRequestCheckModel]{
	Name: "merge:check",
	Run: func(c *cmd.Command[MergeRequestCheckModel], args []string) (MergeRequestCheckModel, error) {
		search := &gitlab.SearchMergeRequest{}

		fs := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fs, search)
		fs.Parse(args)
		projectIds := prIdsDemux.demux(fs.Args())

		if len(projectIds) == 0 {
			return nil, fmt.Errorf("missing projectId")
		}

		// Retrieve all pipelines and jobs async
		type plsj struct {
			projectId   string
			projectName string
			ms          *gitlab.MergeRequestsResponseItemDto
			md          *gitlab.MergeRequestResponseDto
			p           *gitlab.PipelineInfoDto
			r           *gitlab.PipelineJobsResponseDto
			err         error
		}

		// Each goroutines have and internal waitgroup and a channel.
		// Wait all goroutines terminations before close channel.

		// Define channels in pipe
		mergeRequestsC := make(chan *plsj, 10)
		mergeRequestDetailsC := make(chan *plsj, 10)
		pipelinesC := make(chan *plsj, 10)
		jobsC := make(chan *plsj, 10)

		// Find all merge requests by projectsid (close chan mergerequestsc).
		go func() {
			log.Debug("start routine find all merge requests")
			defer close(mergeRequestsC)
			wgi := sync.WaitGroup{}
			for _, projectId := range projectIds {
				projectName := projectId
				projectId := prIds.Get(projectId)
				wgi.Go(func() {
					log.Debug("start routine find merge requests projectId=%s projectName=%q", projectId, projectName)
					mrs, err := gitlabClient.MergeRequests(projectId, search)
					if err != nil {
						mergeRequestsC <- &plsj{
							projectName: projectName,
							projectId:   projectId,
							err:         fmt.Errorf("get pipelinese projectId=%q: %w", projectId, err),
						}
						return
					}
					for _, mr := range mrs.Items {
						mergeRequest := mr
						mergeRequestsC <- &plsj{
							projectName: projectName,
							projectId:   projectId,
							ms:          &mergeRequest,
						}
					}
					log.Debug("end routine find merge requests projectId=%s projectName=%q", projectId, projectName)
				})
			}
			wgi.Wait()
			log.Debug("end routine find all merge requests")
		}()

		// Find all merge request details, range over chan mergeRequestsC, close chan mergeRequestDetailsC.
		go func() {
			log.Debug("start routine getmerge request detail")
			defer close(mergeRequestDetailsC)
			wgi := sync.WaitGroup{}
			for mr := range mergeRequestsC {
				wgi.Go(func() {
					if mr.err != nil {
						mergeRequestDetailsC <- mr
						return
					}
					log.Debug("start routine get merge request detail %d", mr.ms.Id)
					mrId := mr.ms.Iid
					md, err := gitlabClient.MergeRequest(mr.projectId, strconv.Itoa(mrId))
					if err != nil {
						mr.err = fmt.Errorf("get merge request retail projectId=%q, mergeRequestId=%d: %w", mr.projectName, mrId, err)
					} else {
						mr.md = md
					}
					mergeRequestDetailsC <- mr
					log.Debug("done routine get merge request detail %d", mr.ms.Id)
				})
			}
			wgi.Wait()
			log.Debug("end routine getmerge request detail")
		}()

		// Find all pipelines, range over chan mergeRequestDetailsC, close chan pipelinesC.
		go func() {
			log.Debug("start routine get pipeline request detail")
			defer close(pipelinesC)
			wgi := sync.WaitGroup{}
			for row := range mergeRequestDetailsC {
				wgi.Go(func() {
					if row.err != nil {
						pipelinesC <- row
						return
					}
					log.Debug("start routine get pipeline request detail %d", row.md.HeadPipeline.Id)
					row.p = &row.md.HeadPipeline
					pipelinesC <- row
					log.Debug("end routine get pipeline request detail %d", row.md.HeadPipeline.Id)
				})
			}
			wgi.Wait()
			log.Debug("end routine get pipeline request detail")
		}()

		// Find all pipelines, range over chan pipelinesC, close chan jobsC.
		go func() {
			log.Debug("start routine get job request detail")
			defer close(jobsC)
			wgi := sync.WaitGroup{}
			for row := range pipelinesC {
				wgi.Go(func() {
					if row.err != nil {
						jobsC <- row
						return
					}
					log.Debug("start routine get job request detail %d", row.p.Id)
					r, err := gitlabClient.PipelineJobs(row.projectId, strconv.Itoa(row.p.Id), &gitlab.SearchPipelineJob{})
					if err != nil {
						row.err = fmt.Errorf(":%w", err)
					} else {
						row.r = r
					}
					jobsC <- row
					log.Debug("end routine get job request detail %d", row.p.Id)
				})
			}
			wgi.Wait()
			log.Debug("end routine get job request detail")
		}()

		// Map model
		model := make([]MRChPipeline, 0, 50)
		for res := range jobsC { // Channel closed by dedicated goroutine
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
					Id:            res.md.Id,
					Project:       res.projectName,
					PipelineState: res.md.State,
					Pipeline:      res.p.Status,
					MergeStatus:   res.md.MergeStatus,
					TargetBranch:  res.md.TargetBranch,
					UpdatedAt:     res.md.UpdatedAt,
					Jobs:          jobs,
					JobsErrors:    strings.Join(jobses, ","),
					WebUrl:        res.md.WebUrl,
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
	Id            int // pipelineId
	Project       string
	PipelineState string
	Pipeline      string
	MergeStatus   string
	TargetBranch  string
	UpdatedAt     string
	Jobs          string // jobs count ex: 10/12
	JobsErrors    string // jobs error names separated by comma, ex: [job1:job1Id],[job3:job3Id],[job6:job6Id]
	WebUrl        string
}
