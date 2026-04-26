package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

var ConfigPrepareFilesCmd = command{
	Name: "prepare-files",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		files := []string{
			".config/docker/simpl-services/onboarding/keycloak-onboarding-secret-authority.env",
		}
		for _, fp := range files {
			_, err := os.Stat(fp)
			if errors.Is(err, os.ErrNotExist) {
				dir := filepath.Dir(fp)
				if err := os.MkdirAll(dir, 0644); err != nil {
					log.Fatalf("create directory dir %q: %s", dir, err)
				}
				if _, err := os.Create(fp); err != nil {
					log.Fatalf("unable to create file %q: %s", fp, err)
				}
			} else if err != nil {
				log.Fatalf("unable to read file %s: %s", fp, err)
			} else {
				log.Printf("File already exist %q", fp)
			}
		}
		return 0
	},
}
