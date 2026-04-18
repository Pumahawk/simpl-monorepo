package main

var forwardComposeBaseCmd = []string{
	"compose",
	"-f", ".config/docker/forward-node/docker-compose.yaml",
	"-p", "simpl-portforward",
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
	Name: "forward",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		exargs := append(forwardComposeBaseCmd, args...)
		cmd := baseEx("docker", exargs...)
		return runCmd(cmd)
	},
}

var ClusterForwardNodeUpCmd = command{
	Name: "forward-up",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		exargs := []string{"up", "-d"}
		exargs = append(forwardComposeBaseCmd, exargs...)
		exargs = append(exargs, args...)
		cmd := baseEx("docker", exargs...)
		return runCmd(cmd)
	},
}

var ClusterForwardNodeDownCmd = command{
	Name: "forward-down",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		exargs := []string{"down"}
		exargs = append(forwardComposeBaseCmd, exargs...)
		exargs = append(exargs, args...)
		cmd := baseEx("docker", exargs...)
		return runCmd(cmd)
	},
}
