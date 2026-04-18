package kc

import (
	"simplcli/internal/ex"
)

func Wait() error {
	cmd := ex.New(
		"kubectl",
		"wait", "-n", "local-authority",
		"--for=jsonpath={.status.availableReplicas}=1",
		"statefulset.apps/local-authority-keycloak",
	)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
