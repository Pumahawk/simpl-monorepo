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

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

// Autoheal merge requests given multiple project ids.
// Find opened merge request with head piplined failed with fortify error.
// Try to mark false positives on fortify and restart gitlab job.
var GitlabAutoHealMergeCmd = cmd.Command[[]GitlabAutoHealMergeCmdModel]{
	Name: "autoheal:merge",
	Run: func(c *cmd.Command[[]GitlabAutoHealMergeCmdModel], args []string) ([]GitlabAutoHealMergeCmdModel, error) {
		params := &GitlabAutoHealMergeFlags{}
		fl := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fl, params)
		fl.Parse(args)

		projectIds := prIdsDemux.demux(fl.Args())

		if len(projectIds) == 0 {
			return nil, fmt.Errorf("missing project id")
		}

		//  Get opened merge requests given multiple projects
		type mergeRequestsT struct {
			prId   string
			prName string
			r      *gitlab.MergeRequestsResponseItemDto
		}
		mergeRequestsC := make(chan mergeRequestsT, 10)
		go func() {
			defer close(mergeRequestsC)
			wg := &sync.WaitGroup{}
			for _, projectId := range projectIds {
				projectName := projectId
				projectId := prIds.Get(projectId)
				wg.Go(func() {
					search := &gitlab.SearchMergeRequest{
						TargetBranch: params.TargetBranch,
						State:        "opened",
					}
					r, err := gitlabClient.MergeRequests(projectId, search)
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to retrieve merge request projectName=%q: %s\n", projectName, err)
						return
					}
					for i := range r.Items {
						item := &r.Items[i]
						mergeRequestsC <- mergeRequestsT{projectId, projectName, item}
					}
				})
			}
			wg.Wait()
		}()

		// Get merge request details to extract the head pipeline
		type pipelinesT struct {
			prId   string
			prName string
			r      *gitlab.PipelineInfoDto
		}
		pipelinesC := make(chan pipelinesT, 10)
		go func() {
			defer close(pipelinesC)
			wg := &sync.WaitGroup{}
			for mr := range mergeRequestsC {
				prId := mr.prId
				prName := mr.prName
				mrId := strconv.Itoa(mr.r.Iid)
				wg.Go(func() {
					r, err := gitlabClient.MergeRequest(prId, mrId)
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to retrieve merge request details projectId=%q mergeRequestId=%q: %s\n", prName, mrId, err)
						return
					}
					pipelinesC <- pipelinesT{prId, prName, &r.HeadPipeline}
				})
			}
			wg.Wait()
		}()

		// Extraction of all jobs from head pipeline
		type jobsT struct {
			prId   string
			prName string
			r      *gitlab.PipelineJobsResponseItemDto
		}
		jobsC := make(chan jobsT, 10)
		go func() {
			defer close(jobsC)
			wg := &sync.WaitGroup{}
			for pl := range pipelinesC {
				prId := pl.prId
				prName := pl.prName
				plId := strconv.Itoa(pl.r.Id)
				wg.Go(func() {
					r, err := gitlabClient.PipelineJobs(prId, plId, &gitlab.SearchPipelineJob{})
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to retrieve jobs projectId=%q pipelineId=%q: %s\n", prId, plId, err)
						return
					}
					for i := range r.Items {
						job := &r.Items[i]
						jobsC <- jobsT{prId, prName, job}
					}
				})
			}
			wg.Wait()
		}()

		// Retrieve fortify releaseId form job's log
		type fortiFyRefT struct {
			prId   string
			prName string
			jobId  string
			r      string
		}
		fortiFyRefC := make(chan fortiFyRefT, 10)
		go func() {
			defer close(fortiFyRefC)
			wg := &sync.WaitGroup{}
			for jb := range jobsC {
				prId := jb.prId
				prName := jb.prName
				jobId := strconv.Itoa(jb.r.Id)
				if strings.Contains(jb.r.Name, "fortify") {
					wg.Go(func() {
						log, err := gitlabClient.JobLog(prId, jobId)
						if err != nil {
							fmt.Fprintf(os.Stderr, "unable to retrieve job log\n")
						}
						sc := bufio.NewScanner(bytes.NewBuffer(log))
						for sc.Scan() {
							line := sc.Text()
							if strings.Contains(line, "[MASKED]/Redirect/Releases/") {
								words := strings.Split(line, "/")
								releaseId := words[len(words)-1]
								fortiFyRefC <- fortiFyRefT{prId, prName, jobId, releaseId}
							}
						}
					})
				}
			}
			wg.Wait()
		}()

		// Mark false positive on fortify
		type fortifyFPT struct {
			prId   string
			prName string
			jobId  string
			r      string
		}
		fortifyFPC := make(chan fortifyFPT, 10)
		go func() {
			defer close(fortifyFPC)
			wg := &sync.WaitGroup{}
			for release := range fortiFyRefC {
				prId := release.prId
				prName := release.prName
				jobId := release.jobId
				releaseId := release.r
				wg.Go(func() {
					if err := fortifyClient.MarkFalsePositive(releaseId); err != nil {
						fmt.Fprintf(os.Stderr, "unable to mark false positive projectId=%q release=%q: %s\n", prId, releaseId, err)
						return
					}
					fortifyFPC <- fortifyFPT{prId, prName, jobId, releaseId}
				})
			}
			wg.Wait()
		}()

		// Restart fortify in head pipeline job
		type fortifyFinalT struct {
			prId   string
			prName string
			jobId  string
			r      string
		}
		fortifyFinalC := make(chan fortifyFinalT, 10)
		go func() {
			defer close(fortifyFinalC)
			wg := &sync.WaitGroup{}
			for release := range fortifyFPC {
				prId := release.prId
				prName := release.prName
				jobId := release.jobId
				releaseId := release.r
				wg.Go(func() {
					r, err := gitlabClient.JobRetry(prId, jobId)
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to restart projectId=%q jobId=%q: %s\n", prId, jobId, err)
						return
					}
					fortifyFinalC <- fortifyFinalT{prId, prName, strconv.Itoa(r.Id), releaseId}
				})
			}
			wg.Wait()
		}()

		// Generate report
		ret := make([]GitlabAutoHealMergeCmdModel, 0, 10)
		for v := range fortifyFinalC {
			ret = append(ret, GitlabAutoHealMergeCmdModel{
				ProjectName: v.prName,
				ProjectId:   v.prId,
				JobdId:      v.jobId,
				Release:     v.r,
			})
		}

		return ret, nil
	},
}

type GitlabAutoHealMergeCmdModel struct {
	ProjectName string
	ProjectId   string
	JobdId      string
	Release     string
}

type GitlabAutoHealMergeFlags struct {
	TargetBranch string
}
