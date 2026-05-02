package main

import (
	"flag"
	"fmt"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var GitlabRegistryCmd = cmd.Command[any]{
	Name: "reg",
	Run: func(c *cmd.Command[any], args []string) (any, error) {
		search := &gitlab.SearchRegistry{}

		fl := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fl, search)
		fl.Parse(args)
		projectId := fl.Arg(0)

		if projectId == "" {
			return nil, fmt.Errorf("missing project id")
		}

		r, err := gitlabClient.Registry(prIds.Get(projectId), search)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve registry, project %q: %w", projectId, err)
		}

		return r, nil
	},
}
