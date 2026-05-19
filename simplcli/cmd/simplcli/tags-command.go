package main

import (
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/gitlab"
)

type TagItemDto struct {
	Project   string `json:"project"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Message   string `json:"message"`
	Protected bool   `json:"protected"`
	CreatedAt string `json:"created_at"`
}

var GitlabTagsCmd = SearchMultiProjectAsyncCmd[gitlab.SearchTags, struct{}, TagItemDto]{
	Name: "tags:search",
	ApiFunc: func(projectId string, flags struct{}, search *gitlab.SearchTags) ([]TagItemDto, error) {
		r, err := gitlabClient.Tags(projectId, search)
		if err != nil {
			return nil, err
		}
		res := make([]TagItemDto, 0, len(r))
		for _, tag := range r {
			res = append(res, TagItemDto{
				Project:   projectId,
				Name:      tag.Name,
				Target:    tag.Target,
				Message:   tag.Message,
				Protected: tag.Protected,
				CreatedAt: tag.CreatedAt,
			})
		}
		return res, nil
	},
}
