package main

import (
	"fmt"
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
		if err := kc.Wait(1000 * time.Second); err != nil {
			return 1
		}
		if err := kc.Login(); err != nil {
			return 1
		}

		// Get secrets
		authoritySec, err := kc.GetSecret("authority", "cli")
		if err != nil {
			fmt.Printf("error retrieve secrets authority cli: %s\n", err)
			return 1
		}
		onboardingSec, err := kc.GetSecret("onboarding", "onboarding-sa")
		if err != nil {
			fmt.Printf("error retrieve secrets onboarding onboarding-sa: %s\n", err)
			return 1
		}
		testSec, err := kc.GetSecret("authority", "test-client")
		if err != nil {
			fmt.Printf("error retrieve secrets authority test-client: %s\n", err)
			return 1
		}
		consumerSec, err := kc.GetSecret("consumer", "cli")
		if err != nil {
			fmt.Printf("error retrieve secrets consumer cli: %s\n", err)
			return 1
		}

		fmt.Printf("secret authority cli: %q\n", authoritySec)
		fmt.Printf("secret onboarding onboarding-sa: %q\n", onboardingSec)
		fmt.Printf("secret authority test-client: %q\n", testSec)
		fmt.Printf("secret consumer cli: %q\n", consumerSec)

		if err := ex.RunList(
			ex.New("yq", "eval", fmt.Sprintf(".environments[0].service-accounts[0].secret = %q", authoritySec), "-i", testEnvConfPath).Run,
			ex.New("yq", "eval", fmt.Sprintf(".environments[0].keycloak-service-accounts.applicant.secret = %q", onboardingSec), "-i", testEnvConfPath).Run,
			ex.New("yq", "eval", fmt.Sprintf(".environments[0].keycloak-service-accounts.primary.secret = %q", testSec), "-i", testEnvConfPath).Run,
			ex.New("yq", "eval", fmt.Sprintf(".environments[1].service-accounts[].secret = %q", consumerSec), "-i", testEnvConfPath).Run,
		); err != nil {
			return 1
		}
		return 0
	},
}
