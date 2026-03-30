package main

type RunCmd = func([]string)

type Command struct {
	Name        string
	Description string
	Run         RunCmd
}
