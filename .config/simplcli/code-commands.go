package main

import (
	"path"
)

var pathBackend = path.Join("simpl-repo", "backend")

var CodeBuildNoTestCmd = command{
	Name: "build:no-test",
	Descr: "" +
		"",
	Func: func(args ...string) int8 {
		mvndargs := []string{
			"-Dmvnd.connectTimeout=100000", "clean", "install", "-fae", "-f", pathBackend,
			"--fail-never", "-Dlicense.skipDownloadLicenses", "-Dmaven.test.skip",
		}
		mvndargs = append(mvndargs, args...)
		cmd := baseEx("mvnd", mvndargs...)
		if err := cmd.Run(); err != nil {
			return int8(cmd.ProcessState.ExitCode())
		} else {
			return 0
		}
	},
}
