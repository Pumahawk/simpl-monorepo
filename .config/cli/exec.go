package main

import (
	"fmt"
	"os"
	"os/exec"
)

func KeycloakConfigCliExec(dir, realm string) error {
	fmt.Printf("Start keycloak-config-cli realm=[%s]\n", realm)
	cmd := exec.Command("java",
		"-jar", "bin/keycloak-config-cli.jar",
		"--keycloak.url=http://localhost:8100/auth",
		"--keycloak.user=admin",
		"--keycloak.password=admin",
		"--import.files.locations=out/realm.yaml",
	)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("keycloak config cli dir=[%s] realm=[%s]: %v", dir, realm, err)
	}
	return nil
}

func GomplateExec(dir, realm string) error {
	fmt.Printf("Execute gomplate dir=[%s] realm=[%s]", dir, realm)
	cmd := exec.Command("gomplate", "-c", fmt.Sprintf(".=examples/%s-with-eidas.yaml", realm))
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gomplate exec error dir=[%s] realm=[%s]: %w", dir, realm, err)
	}
	return nil
}
