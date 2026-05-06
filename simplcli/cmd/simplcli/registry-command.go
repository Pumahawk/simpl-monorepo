package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/Pumahawk/simpl-monorepo/internal/cmd"
	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

type GitlabRegistryItemModel struct {
	Id          int    `json:"id"`
	Project     string `json:"project"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	PackageType string `json:"package_type"`
	CreatedAt   string `json:"created_at"`
	CreatorId   int    `json:"creator_id"`
}

var GitlabRegistryCmd = cmd.Command[[]GitlabRegistryItemModel]{
	Name: "registry:search",
	Run: func(c *cmd.Command[[]GitlabRegistryItemModel], args []string) ([]GitlabRegistryItemModel, error) {
		search := &gitlab.SearchRegistry{}

		fl := flag.NewFlagSet("", flag.ExitOnError)
		structFlag(fl, search)
		fl.Parse(args)
		projectIds := prIdsDemux.demux(fl.Args())

		if len(projectIds) == 0 {
			return nil, fmt.Errorf("missing project id")
		}

		ch := make(chan *GitlabRegistryItemModel)
		go func() {
			defer close(ch)
			wg := &sync.WaitGroup{}
			wg.Go(func() {
				for i := range projectIds {
					projectName := projectIds[i]
					projectId := prIds.Get(projectName)
					r, err := gitlabClient.Registry(prIds.Get(projectId), search)
					if err != nil {
						fmt.Fprintf(os.Stderr, "unable to retrieve registry, project %q: %s\n", projectId, err)
						return
					}
					for i := range r.Items {
						item := &r.Items[i]
						ch <- &GitlabRegistryItemModel{
							Id:          item.Id,
							Project:     projectName,
							Name:        item.Name,
							Version:     item.Version,
							PackageType: item.PackageType,
							CreatedAt:   item.CreatedAt,
							CreatorId:   item.CreatorId,
						}
					}
				}
			})
			wg.Wait()
		}()

		items := make([]GitlabRegistryItemModel, 0, 20*len(prIds))
		for item := range ch {
			items = append(items, *item)
		}

		return items, nil
	},
}
