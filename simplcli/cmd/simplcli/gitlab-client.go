package main

import "github.com/Pumahawk/simpl-monorepo/internal/gitlab"

var gitlabClient = gitlab.Client{
	BaseUrl: "https://code.europa.eu/api/v4",
}
