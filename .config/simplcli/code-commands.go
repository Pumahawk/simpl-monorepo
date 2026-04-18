package main

import (
	"fmt"
	"log"
	"path"
	"simplcli/internal/ex"
	"simplcli/internal/kc"
	"time"
)

var pathBackend = path.Join("simpl-repo", "backend")
var testEnvConfPath = path.Join("simpl-repo", "backend", "tests", "test-automation", "src", "test", "resources", "environments", "environments-local.yaml")

var CodeBuildNoTestCmd = command{
	Name: "build:no-test",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		mvndargs := []string{
			"-Dmvnd.connectTimeout=100000", "clean", "install", "-fae", "-f", pathBackend,
			"--fail-never", "-Dlicense.skipDownloadLicenses", "-Dmaven.test.skip",
		}
		mvndargs = append(mvndargs, args...)
		cmd := ex.New("mvnd", mvndargs...)
		if err := cmd.Run(); err != nil {
			return int8(cmd.ProcessState.ExitCode())
		} else {
			return 0
		}
	},
}

var CodeTestAutomationAutoconfigureClientSecretCmd = command{
	Name: "test-automation:autoconfigure-client-secret",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {

		// Wait keycloak and login
		log.Println("waiting keycloak startup")
		if err := kc.Wait(1000 * time.Second); err != nil {
			return 1
		}
		if err := kc.Login(); err != nil {
			return 1
		}

		// Get secrets from keycloak
		type scr struct{ realm, id, sec, dest string }
		scrs := []*scr{
			{"authority", "cli", "", ".environments[0].service-accounts[0].secret"},
			{"onboarding", "onboarding-sa", "", ".environments[0].keycloak-service-accounts.applicant.secret"},
			{"authority", "test-client", "", ".environments[0].keycloak-service-accounts.primary.secret"},
			{"consumer", "cli", "", ".environments[1].service-accounts[].secret"},
		}
		for _, sc := range scrs {
			if r, err := kc.GetSecret(sc.realm, sc.id); err != nil {
				log.Printf("error retrieve secrets %s %s: %s", sc.realm, sc.id, err)
				return 1
			} else {
				sc.sec = r
			}
		}

		for _, sc := range scrs {
			log.Printf("secret %s %s: %q", sc.realm, sc.id, sc.sec)
		}

		// Write all secrets to test configuration  environments-local.yaml
		runls := make([]func() error, 0, len(scrs))
		for _, sc := range scrs {
			cmd := ex.New("yq", "eval", fmt.Sprintf("%s = %q", sc.dest, sc.sec), "-i", testEnvConfPath)
			runls = append(runls, cmd.Run)
		}
		if err := ex.RunList(runls...); err != nil {
			return 1
		}
		return 0
	},
}
