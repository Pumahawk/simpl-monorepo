package main

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/cmd"
)

var FortifyFalsePositiveCmd = cmd.Command[int]{
	Name: "false-positive",
	Run: func(c *cmd.Command[int], args []string) (int, error) {
		var releaseId string
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.Parse(args)
		releaseId = fs.Arg(0)

		if releaseId == "" {
			return 1, fmt.Errorf("missing argument releaseId")
		}

		if err := fortifyClient.MarkFalsePositive(releaseId); err != nil {
			return 1, fmt.Errorf("fortify mark false positive releaseId=%q", releaseId)
		}
		return 0, nil
	},
}
