package main

import (
	"sort"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
)

type info struct {
	Name string
	Id   string
}

var InfoProjectIdsCmd = cmd.Command[[]info]{
	Name: "projects",
	Run: func(c *cmd.Command[[]info], args []string) ([]info, error) {
		infos := make([]info, 0, len(prIds))
		for k, v := range prIds {
			infos = append(infos, info{k, v})
		}
		sort.Slice(infos, func(i, j int) bool {
			return infos[i].Id < infos[j].Id
		})
		return infos, nil
	},
}
