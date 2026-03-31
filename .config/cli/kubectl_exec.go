package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func KubectlGetExec(v any, args ...string) error {
	karg := []string{
		"-n", "local-authority",
		"get",
		"-o", "json"}
	karg = append(karg, args...)
	cmd := exec.Command("kubectl", karg...)
	cmd.Stderr = os.Stderr
	var out io.ReadCloser
	var err error
	if out, err = cmd.StdoutPipe(); err != nil {
		return fmt.Errorf("unable to create stdoutpipe args=[%v]: %w", args, err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("unable to start kubectl %v: %w", karg, err)
	}
	if err := json.NewDecoder(out).Decode(v); err != nil {
		return fmt.Errorf("unable to decode kubectl execution: %w", err)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error on exit from kubectl %v: %w", karg, err)
	}
	return nil
}

type KItems struct {
	Items []KItem
}

type KItem struct {
	Metadata KMetadata
}

type KMetadata struct {
	Name string
}
