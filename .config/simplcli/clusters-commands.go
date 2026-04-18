package main

var ClusterCreateCmd = command{
	Name:  "create",
	Descr: "Starts Minikube with a dedicated Docker network and profile.",
	Func: func(args ...string) int8 {
		cmd := baseEx("minikube", "start", "--namespace", "local-authority", "--network=simpl-network", "--driver=docker", "--profile=simpl-control-plane")
		return runCmd(cmd)
	},
}
