package main

import (
	"flag"
	"fmt"
	"os"
)

var JobLogCmd = Command[int]{
	Name: "jobl",
	Run: func(c *Command[int], args []string) (int, error) {

		fl := flag.NewFlagSet("", flag.ExitOnError)
		fl.Parse(args)

		projectId := fl.Arg(0)
		jobId := fl.Arg(1)

		if projectId == "" {
			return 1, fmt.Errorf("missing project id")
		}

		if jobId == "" {
			return 1, fmt.Errorf("missing job id")
		}

		log, err := gitlabClient.JobLog(prIds.Get(projectId), jobId)
		if err != nil {
			return 1, err
		}

		_, err = os.Stdout.Write(log)
		return 0, nil
	},
}
