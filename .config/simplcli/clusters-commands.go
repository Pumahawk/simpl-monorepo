package main

import (
	"path"
	"simplcli/internal/ex"
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
