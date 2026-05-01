package main

import (
	"os"

	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var gitlabClient = gitlab.Client{
	BaseUrl: "https://code.europa.eu/api/v4",
	TokenFunc: func() (string, error) {
		t, _ := os.LookupEnv("GITLAB_TOKEN")
		return t, nil
	},
}
