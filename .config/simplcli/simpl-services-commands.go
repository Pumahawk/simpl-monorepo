package main

import (
	dkc "simplcli/internal/docker-compose"
	"simplcli/internal/ex"
)

var dkCmpSimplServ = dkc.Conf{
	Profile: "simpl-services",
	File:    ".config/docker/simpl-services/docker-compose.yaml",
	EnvFile: "./simpl-services.env",
}

var dkCmpSimplFeServ = dkc.Conf{
	Profile: "simpl-services-fe",
	File:    ".config/docker/simpl-fe/docker-compose.yaml",
	EnvFile: "./simpl-services.env",
}

var SimplServicesTier1AuthorityUpCmd = command{
	Name:  "tier1-authority-up",
	Descr: "",
	Func: func(args ...string) int8 {
		cmd := dkCmpSimplServ.Cmd("up", "-d", "tier1-gateway-authority")
		return ex.RunCmd(cmd)
	},
}

var SimplServicesUpCmd = command{
	Name:  "up",
	Descr: "",
	Func: func(args ...string) int8 {
		cmd := dkCmpSimplServ.Cmd("up", "-d", "--pull=always")
		return ex.RunCmd(cmd)
	},
}

var SimplServicesFeUpCmd = command{
	Name:  "up",
	Descr: "",
	Func: func(args ...string) int8 {
		cmd := dkCmpSimplFeServ.Cmd("up", "-d", "--pull=always")
		return ex.RunCmd(cmd)
	},
}
