package main

import "os"

var pipelinesCmds = CommandGroup("",
	CommandGroup("gl",
		GitlabPipelinesCmd,
		GitlabPipelineCmd,
	),
)

func main() {
	os.Exit(pipelinesCmds.Run(&pipelinesCmds, os.Args[1:]))
}
