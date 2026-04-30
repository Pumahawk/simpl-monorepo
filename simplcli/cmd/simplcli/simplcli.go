package main

import "os"

var pipelinesCmds = CommandGroup("",
	CommandGroup("gl", GitlabPipelinesCmd),
)

func main() {
	os.Exit(pipelinesCmds.Run(&pipelinesCmds, os.Args[1:]))
}
