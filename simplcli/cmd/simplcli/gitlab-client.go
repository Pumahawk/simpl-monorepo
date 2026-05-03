package main

import (
	"fmt"
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
	TokenFunc: func() (*fortify.AuthData, error) {
		envtu := "FORTIFY_USER"
		envtk := "FORTIFY_TOKEN"
		u, _ := os.LookupEnv(envtu)
		t, _ := os.LookupEnv(envtk)
		if u == "" || t == "" {
			return nil, fmt.Errorf("not found valid auth configuration, define env %q %q", envtu, envtk)
		}
		return &fortify.AuthData{Token: t, User: u}, nil
	},
}
