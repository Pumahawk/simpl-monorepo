package main

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

var ClusterCreateCmd = command{
	Name:  "create",
	Descr: "Starts Minikube with a dedicated Docker network and profile.",
	Func: func(args ...string) int8 {
		cmd := baseEx("minikube", "start", "--namespace", "local-authority", "--network=simpl-network", "--driver=docker", "--profile=simpl-control-plane")
		return runCmd(cmd)
	},
}

var ClusterForwardNodeComposeCmd = command{
	Name:  "forward",
	Descr: "",
	Func: func(args ...string) int8 {
		exargs := append(forwardComposeBaseCmd, args...)
		cmd := baseEx("docker", exargs...)
		return runCmd(cmd)
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
		cmd := baseEx("docker", exargs...)
		return runCmd(cmd)
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
