package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"simplcli/internal/ejbca"
	"simplcli/internal/ex"
	"simplcli/internal/kube"
	"time"
)

var ClusterInitializationEjbcaCmd = command{
	Name: "initialization_ejbca",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		var dockerDir = ".config/docker/simpl-services"
		var ejbcaDockerDir = path.Join(dockerDir, "ejbca")
		var identityDockerDir = path.Join(dockerDir, "identity-provider")
		var remotePodConfigDir = "/tmp/config"

		var ej *kube.Item

	loop:
		for {
			ej, err := ejbca.Get()
			if err != nil && err != ejbca.NotFound {
				log.Printf("Unable to get ejbca pod: %s", err)
				return 1
			}
			if err == ejbca.NotFound {
				log.Printf("Not found ejbca pod")
			} else {
				for _, c := range ej.Status.ContainerStatuses {
					if c.Name == "ejbca" {
						if c.Ready {
							break loop
						} else {
							log.Printf("ejbca not ready")
						}
					}
				}
			}
			n := time.Duration(10)
			log.Printf("wait %d seconds", n)
			<-time.After(n * time.Second)
		}
		log.Printf("Ejbca ready. Start initialization")
		pn := ej.Metadata.Name
		cmdMkdir := ex.New("kubectl", "exec", "-n", "local-authority", pn, "--", "mkdir", remotePodConfigDir)
		if err := cmdMkdir.Run(); err != nil {
			fmt.Printf("unable to create directory %q, pod=%q: %s", remotePodConfigDir, pn, err)
			return 1
		}
		files, err := os.ReadDir(ejbcaDockerDir)
		if err != nil {
			fmt.Printf("unable to read dir %q", ejbcaDockerDir)
			return 1
		}

		panic("not implemented")
	},
}
var ClusterDownloadEjbcaPemCmd = command{
	Name: "cluster:download_ejbca_pem",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var ConfigPrepareFilesCmd = command{
	Name: "config:prepare-files",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var SimplServicesTier1AuthorityUpCmd = command{
	Name: "simpl-services:tier1-authority-up",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var ClusterKeycloakAutoconfigureCmd = command{
	Name: "cluster:keycloak-autoconfigure",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var SimplServicesUpCmd = command{
	Name: "simpl-services:up",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var SimplServicesFeUpCmd = command{
	Name: "simpl-services-fe:up",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}

var ClusterAuthorityInitializationCmd = command{
	Name: "cluster:authority_initialization",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		panic("not implemented")
	},
}
