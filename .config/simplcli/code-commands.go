package main

import (
	"fmt"
	"path"
	"simplcli/internal/ex"
	"simplcli/internal/kc"
	"time"
)

var pathBackend = path.Join("simpl-repo", "backend")

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
			fmt.Printf("error retrieve secrets authoritySec: %s\n", err)
			return 1
		}
		onboardingSec, err := kc.GetSecret("onboarding", "onboarding-sa")
		if err != nil {
			fmt.Printf("error retrieve secrets onboardingSec: %s\n", err)
			return 1
		}
		testSec, err := kc.GetSecret("authority", "test-client")
		if err != nil {
			fmt.Printf("error retrieve secrets testSec: %s\n", err)
			return 1
		}
		consumerSec, err := kc.GetSecret("consumer", "cli")
		if err != nil {
			fmt.Printf("error retrieve secrets consumerSec: %s\n", err)
			return 1
		}

		// Print secrets, TODO write secrets to file
		fmt.Printf("secret authoritySec: %q\n", authoritySec)
		fmt.Printf("secret onboardingSec: %q\n", onboardingSec)
		fmt.Printf("secret testSec: %q\n", testSec)
		fmt.Printf("secret consumerSec: %q\n", consumerSec)
		return 0
	},
}
