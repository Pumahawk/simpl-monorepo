package main

import (
	"os"

	"github.com/Pumahawk/simpl-monorepo/internal/fortify"
	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

var gitlabClient = gitlab.Client{
	BaseUrl: "https://code.europa.eu",
	TokenFunc: func() (string, error) {
		t, _ := os.LookupEnv("GITLAB_TOKEN")
		return t, nil
	},
}

var fortifyClient = fortify.Client{
	BaseUrl: "https://api.emea.fortify.com",
	TokenFunc: func() (string, error) {
		t, _ := os.LookupEnv("FORTIFY_TOKEN")
		return t, nil
	},
}
