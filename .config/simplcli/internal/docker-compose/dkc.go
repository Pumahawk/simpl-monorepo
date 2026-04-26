package dkc

import (
	"os/exec"
	"simplcli/internal/ex"
)

type Conf struct {
	Profile string
	File    string
	EnvFile string
}

func (c *Conf) Cmd(args ...string) *exec.Cmd {
	cargs := []string{"compose"}
	if c.Profile != "" {
		cargs = append(cargs, "-p", c.Profile)
	}
	if c.File != "" {
		cargs = append(cargs, "-f", c.File)
	}
	if c.EnvFile != "" {
		cargs = append(cargs, "--env-file="+c.EnvFile)
	}
	cargs = append(cargs, args...)
	return ex.New("docker", cargs...)
}
