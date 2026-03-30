package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

var commands = []Command{
	KeycloakRealmAutoconfigure,
}

func main() {
	if len(os.Args) > 1 {
		cname := os.Args[1]
		cargs := os.Args[2:]
		for _, command := range commands {
			if command.Name == cname {
				command.Run(cargs)
				return
			}
		}
		fmt.Printf("Unknow subcommand: %s\n\n", cname)
	}
	PrintHelp()
}

func PrintHelp() {
	t := template.Must(template.New("help-message").
		Parse(`subcommand:
{{- range .Commands }}
- {{ .Name }}: {{ .Description }}
{{- end }}
- help: Print this message
`))
	v := struct {
		Commands []Command
	}{
		Commands: commands,
	}
	t.Execute(os.Stdout, v)
}

type RunCmd = func([]string)

type Command struct {
	Name        string
	Description string
	Run         RunCmd
}

const KeycloakRealmAutoconfigureName = "realm-autoconfigure"
const KeycloakRealmConfiguratorDirectory = "simpl-repo/backend/misc/tier1-authentication/tools/keycloak-realm-configurator"

var KeycloakRealmAutoconfigure Command = Command{
	Name:        KeycloakRealmAutoconfigureName,
	Description: "keycloak execution",
	Run: func(args []string) {
		fl := flag.NewFlagSet("", flag.ExitOnError)
		kconfd := fl.String(KeycloakRealmAutoconfigureName, KeycloakRealmConfiguratorDirectory, "Path to keycloak realm configurator")
		fl.Parse(args)
		var realms = [...]string{"authority", "participant", "onboarding"}
		fmt.Println("Generate Realm")
		for _, realm := range realms {
			fmt.Printf("process realm=[%s]\n", realm)
			cmd := exec.Command("gomplate", "-c", fmt.Sprintf(".=examples/%s-with-eidas.yaml", realm))
			cmd.Dir = *kconfd
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("keycloak realm autoconfigure: import realm %s error: %v\n", realm, err)
				os.Exit(1)
			}
			for i := range 3 {
				fmt.Printf("Try (%d/%d)\n", i+1, 3)
				if err := KeycloakConfigCliExec(*kconfd, realm); err != nil {
					fmt.Printf("Error in keycloak configuration. dir=[%s] realm=[%s]. Retry.\n", *kconfd, realm)
					if 3 == i+1 {
						fmt.Printf("Failed max retry. Stop program. Error: %v\n", err)
						os.Exit(1)
					}
				} else {
					fmt.Printf("Success autoconfiguration realm [%s]\n", realm)
					break
				}
			}
			fmt.Printf("Success autconfigure realm=[%s]\n", realm)
		}
	},
}

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
