package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"simplcli/internal/ejbca"
	"simplcli/internal/ex"
	"simplcli/internal/kube"
	"time"
)

const knamespace = "local-authority"

var forwardComposeBaseCmd = []string{
	"compose",
	"-f", ".config/docker/forward-node/docker-compose.yaml",
	"-p", "simpl-portforward",
}

var redPandaComposeBaseCmd = []string{
	"compose",
	"-f", ".config/docker/redpanda/docker-compose.yaml",
	"-p", "simpl-redpanda",
}

var chartLocalAuthorityPath = path.Join("simpl-repo", "docs", "charts", "app-values", "local", "local-authority")

var ClusterCreateCmd = command{
	Name:  "create",
	Descr: "Starts Minikube with a dedicated Docker network and profile.",
	Func: func(args ...string) int8 {
		cmd := ex.New("minikube", "start", "--namespace", knamespace, "--network=simpl-network", "--driver=docker", "--profile=simpl-control-plane")
		return ex.RunCmd(cmd)
	},
}

var ClusterForwardNodeComposeCmd = command{
	Name:  "forward",
	Descr: "",
	Func: func(args ...string) int8 {
		exargs := append(forwardComposeBaseCmd, args...)
		cmd := ex.New("docker", exargs...)
		return ex.RunCmd(cmd)
	},
}

var ClusterForwardNodeUpCmd = command{
	Name:  "forward-up",
	Descr: "",
	Func: func(args ...string) int8 {
		args = append(args, "up", "-d")
		return ClusterForwardNodeComposeCmd.Func(args...)
	},
}

var ClusterForwardNodeDownCmd = command{
	Name:  "forward-down",
	Descr: "",
	Func: func(args ...string) int8 {
		args = append(args, "down")
		return ClusterForwardNodeComposeCmd.Func(args...)
	},
}

var ClusterRedpandaComposeCmd = command{
	Name:  "redpanda",
	Descr: "",
	Func: func(args ...string) int8 {
		exargs := append(redPandaComposeBaseCmd, args...)
		cmd := ex.New("docker", exargs...)
		return ex.RunCmd(cmd)
	},
}

var ClusterRedpandaUpCmd = command{
	Name:  "redpanda:up",
	Descr: "",
	Func: func(args ...string) int8 {
		args = append(args, "up", "-d")
		return ClusterRedpandaComposeCmd.Func(args...)
	},
}

var ClusterRedpandaDownCmd = command{
	Name:  "redpanda:down",
	Descr: "",
	Func: func(args ...string) int8 {
		args = append(args, "down")
		return ClusterRedpandaComposeCmd.Func(args...)
	},
}

var ClusterAuthorityInstallOrUpgradeCmd = command{
	Name:  "authority-install-or-upgrade",
	Descr: "",
	Func: func(args ...string) int8 {
		cmdHelm := ex.New("helm",
			"upgrade",
			"-i",
			"--dependency-update",
			"--create-namespace",
			"-n", knamespace,
			"--set", "keycloak.livenessProbe.enabled=false",
			"--set-string", "postgrest.primary.initdb.scripts.restoreEjbca=",
			knamespace,
			chartLocalAuthorityPath,
		)
		cmdKubens := ex.New("kubens", knamespace)
		rps0 := ex.New("kubectl", "-n", knamespace, "scale", "--replicas", "0", "deployment.apps/redpanda")
		rps1 := ex.New("kubectl", "-n", knamespace, "scale", "--replicas", "0", "deployment.apps/redpanda-console")
		if err := ex.RunList(
			cmdHelm.Run,
			cmdKubens.Run,
			rps0.Run,
			rps1.Run,
		); err != nil {
			return 1
		}
		return 0
	},
}

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
			ejpod, err := ejbca.Get()
			if err != nil && err != ejbca.NotFound {
				log.Printf("Unable to get ejbca pod: %s", err)
				return 1
			}
			if err == ejbca.NotFound {
				log.Printf("Not found ejbca pod")
			} else {
				for _, c := range ejpod.Status.ContainerStatuses {
					if c.Name == "ejbca" {
						if c.Ready {
							ej = ejpod
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
		cmdMkdir := ex.New("kubectl", "exec", "-n", "local-authority", pn, "--", "mkdir", "-p", remotePodConfigDir)
		if err := cmdMkdir.Run(); err != nil {
			fmt.Printf("unable to create directory %q, pod=%q: %s\n", remotePodConfigDir, pn, err)
			return 1
		}
		files, err := os.ReadDir(path.Join(ejbcaDockerDir, "config"))
		if err != nil {
			fmt.Printf("unable to read dir %q\n", ejbcaDockerDir)
			return 1
		}
		for _, f := range files {
			path := path.Join(ejbcaDockerDir, "config", f.Name())
			fo, err := os.Open(path)
			if err != nil {
				fmt.Printf("unable to open file %q: %s\n", path, err)
				return 1
			}
			cpCmd := ex.New("kubectl", "exec", "-i", pn, "--", "sh", "-c", "cat > "+remotePodConfigDir+"/"+f.Name())
			cpCmd.Stdin = fo
			if err := cpCmd.Run(); err != nil {
				fmt.Printf("unable to copy file %q to pod %q\n", path, pn)
				return 1
			}
		}
		entryPointShPath := path.Join(ejbcaDockerDir, "entrypoint.sh")
		entryPointSh, err := os.Open(entryPointShPath)
		if err != nil {
			fmt.Printf("Unable open file %q\n", entryPointShPath)
			return 1
		}

		cpEntryPCmd := ex.New("kubectl", "exec", "-i", pn, "--", "sh", "-")
		cpEntryPCmd.Stdin = entryPointSh
		if err := cpEntryPCmd.Run(); err != nil {
			fmt.Printf("unablet to copy %q to pod %q: %s\n", entryPointShPath, pn, err)
			return 1
		}

		for _, fname := range []string{"SuperAdmin.p12", "truststore.jks"} {
			bf := &bytes.Buffer{}
			cpCmd := ex.New("kubectl", "exec", "-i", pn, "--", "cat", "/opt/keyfactor/p12/"+fname)
			cpCmd.Stdout = bf
			if err := cpCmd.Run(); err != nil {
				fmt.Printf("%s\n", err)
				return 1
			}
			out, err := os.OpenFile(path.Join(identityDockerDir, fname), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("%s\n", err)
				return 1
			}
			if _, err := io.Copy(out, bf); err != nil {
				fmt.Printf("%s\n", err)
				return 1
			}
		}
		return 0
	},
}
