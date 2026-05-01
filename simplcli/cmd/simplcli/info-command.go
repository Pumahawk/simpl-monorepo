package main

import (
	"sort"
)

type info struct {
	Name string
	Id   string
}

var InfoProjectIdsCmd = Command[[]info]{
	Name: "pr",
	Run: func(c *Command[[]info], args []string) ([]info, error) {
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
