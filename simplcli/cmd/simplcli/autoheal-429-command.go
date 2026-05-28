package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/gitlab"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/log"
)

type GitlabAutoHeal429Model struct {
	Project    string
	PipelineId string
	JobId      string
	WebUrl     string
}

var GitlabAutoHeal429Cmd = cmd.Command[[]GitlabAutoHeal429Model]{
	Name: "autoheal:maven-429",
	Run: func(c *cmd.Command[[]GitlabAutoHeal429Model], args []string) ([]GitlabAutoHeal429Model, error) {
		var since string
		var dryRun bool
		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.BoolVar(&dryRun, "dry-run", false, "")
		fl.StringVar(&since, "since", "", "")
		fl.Parse(args)

		projectIds := prIdsDemux.demux(fl.Args())

		if len(projectIds) == 0 {
			return nil, fmt.Errorf("missing project id")
		}

		type newPipe struct{ prId, prName, pipeId string }
		type newJob struct{ prId, prName, pipeId, jobId, jobWebUrl string }
		newPipeC := make(chan newPipe)
		newJobC := make(chan newJob)
		modelC := make(chan GitlabAutoHeal429Model)
		go func() {
			defer close(newPipeC)
			wg := &sync.WaitGroup{}
			for _, prName := range projectIds {
				wg.Go(func() {
					log.Debug("search pipelines projectId=%q", prName)
					pr := prIds.Get(prName)
					r, err := gitlabClient.Pipelines(pr, &gitlab.SearchPipeline{Status: "failed", CreatedAfter: since})
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to search pipelines: %s\n", err)
						return
					}
					for _, p := range r.Items {
						newPipeC <- newPipe{pr, prName, strconv.Itoa(p.Id)}
					}
				})
			}
			wg.Wait()
		}()
		go func() {
			defer close(newJobC)
			wg := &sync.WaitGroup{}
			for ps := range newPipeC {
				wg.Go(func() {
					log.Debug("find pipeline=%q", ps.pipeId)
					jobs, err := gitlabClient.PipelineJobs(ps.prId, ps.pipeId, &gitlab.SearchPipelineJob{})
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to read pipelines jobs: %s\n", err)
						return
					}
					for _, j := range jobs.Items {
						if j.Status == "failed" {
							newJobC <- newJob{ps.prId, ps.prName, ps.pipeId, strconv.Itoa(j.Id), j.WebUrl}
						}
					}
				})
			}
			wg.Wait()
		}()
		go func() {
			defer close(modelC)
			wg := &sync.WaitGroup{}
			for job := range newJobC {
				wg.Go(func() {
					log.Debug("find job pipelineId=%q id=%q", job.pipeId, job.jobId)
					l, err := gitlabClient.JobLog(job.prId, job.jobId)
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to read job logs: %s\n", err)
						return
					}
					scan := bufio.NewScanner(bytes.NewBuffer(l))
					for scan.Scan() {
						line := scan.Text()
						if strings.Contains(line, "429, reason phrase: Too Many Requests") {
							if !dryRun {
								if _, err := gitlabClient.JobRetry(job.prId, job.jobId); err != nil {
									fmt.Fprintf(os.Stderr, "unable to retry job: %s\n", err)
									return
								}
								log.Debug("restart job pipelineId=%q jobId=%q\n", job.pipeId, job.jobId)
							} else {
								log.Debug("dry-run restart job pipelineId=%q jobId=%q\n", job.pipeId, job.jobId)
							}
							modelC <- GitlabAutoHeal429Model{
								Project:    job.prName,
								PipelineId: job.pipeId,
								JobId:      job.jobId,
								WebUrl:     job.jobWebUrl,
							}
							return
						}
					}
				})
			}
			wg.Wait()
		}()
		var out []GitlabAutoHeal429Model
		for m := range modelC {
			out = append(out, m)
		}

		return out, nil
	},
}
