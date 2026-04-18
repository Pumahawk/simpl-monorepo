package main

import (
	"os"
	"os/exec"
)

func runExList(runs ...func() error) error {
	for _, f := range runs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func baseEx(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd
}

func runCmd(c *exec.Cmd) int8 {
	if err := c.Run(); err != nil {
		return int8(c.ProcessState.ExitCode())
	} else {
		return 0
	}
}
