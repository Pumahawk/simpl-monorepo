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
		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.StringVar(&since, "since", "", "")
		fl.Parse(args)

		projectIds := prIdsDemux.demux(fl.Args())

		if len(projectIds) == 0 {
			return nil, fmt.Errorf("missing project id")
		}

		type newPipe struct{ prId, pipeId string }
		type newJob struct{ prId, pipeId, jobId, jobWebUrl string }
		newPipeC := make(chan newPipe)
		newJobC := make(chan newJob)
		go func() {
			defer close(newPipeC)
			wg := &sync.WaitGroup{}
			for _, pr := range projectIds {
				wg.Go(func() {
					log.Debug("search pipelines projectId=%q", pr)
					pr = prIds.Get(pr)
					r, err := gitlabClient.Pipelines(pr, &gitlab.SearchPipeline{Status: "failed", CreatedAfter: since})
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to search pipelines: %s\n", err)
						return
					}
					for _, p := range r.Items {
						newPipeC <- newPipe{pr, strconv.Itoa(p.Id)}
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
							newJobC <- newJob{ps.prId, ps.pipeId, strconv.Itoa(j.Id), j.WebUrl}
						}
					}
				})
			}
			wg.Wait()
		}()
		var out []GitlabAutoHeal429Model
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
						if _, err := gitlabClient.JobRetry(job.prId, job.jobId); err != nil {
							fmt.Fprintf(os.Stderr, "unable to retry job: %s\n", err)
						} else {
							log.Debug("restart job pipelineId=%q jobId=%q\n", job.pipeId, job.jobId)
							out = append(out, GitlabAutoHeal429Model{
								Project:    job.prId,
								PipelineId: job.pipeId,
								JobId:      job.jobId,
								WebUrl:     job.jobWebUrl,
							})
						}
					}
				}
			})
		}
		wg.Wait()
		return out, nil
	},
}
