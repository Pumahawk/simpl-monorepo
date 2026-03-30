package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const KeycloakRealmAutoconfigureName = "realm-autoconfigure"
const KeycloakRealmConfiguratorDirectory = "simpl-repo/backend/misc/tier1-authentication/tools/keycloak-realm-configurator"

var KeycloakRealmAutoconfigure Command = Command{
	Name:        KeycloakRealmAutoconfigureName,
	Description: "keycloak execution",
	Run: func(args []string) {
		fl := flag.NewFlagSet("", flag.ExitOnError)
		kconfd := fl.String(KeycloakRealmAutoconfigureName, KeycloakRealmConfiguratorDirectory, "Path to keycloak realm configurator")
		fl.Parse(args)
		if err := KeycloakRealmAutoconfigureService(*kconfd); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	},
}

func KeycloakRealmAutoconfigureService(kconfd string) error {
	var realms = [...]string{"authority", "participant", "onboarding"}
	fmt.Println("Generate Realm")
	for _, realm := range realms {
		fmt.Printf("process realm=[%s]\n", realm)
		if err := GomplateExec(kconfd, realm); err != nil {
			return fmt.Errorf("keycloak realm autoconfigure: import realm %s error: %w", realm, err)
		}
		for i := range 3 {
			fmt.Printf("Try (%d/%d)\n", i+1, 3)
			if err := KeycloakConfigCliExec(kconfd, realm); err != nil {
				fmt.Printf("Error in keycloak configuration. dir=[%s] realm=[%s]. Retry.\n", kconfd, realm)
				if 3 == i+1 {
					return fmt.Errorf("Failed max retry. Stop program. Error: %w", err)
				}
			} else {
				fmt.Printf("Success autoconfiguration realm [%s]\n", realm)
				break
			}
		}
		fmt.Printf("Success autconfigure realm=[%s]\n", realm)
	}
	return nil
}
