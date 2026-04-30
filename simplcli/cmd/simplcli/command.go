package main

import (
	"fmt"
	"os"
)

type CommandFunc func(c *Command, args []string) int

type Command struct {
	Name string
	Run  CommandFunc
}

func CommandGroup(name string, commands ...Command) Command {
	return Command{
		Name: name,
		Run: func(c *Command, args []string) int {
			if len(args) > 0 {
				for _, c := range commands {
					if c.Name == args[0] {
						return c.Run(&c, args[1:])
					}
				}
				fmt.Fprintf(os.Stdout, "command %q not found\n", args[0])
				return 1
			} else {
				fmt.Printf("commands:\n")
				for _, c := range commands {
					fmt.Printf("%s\n", c.Name)
				}
				return 1
			}
		},
	}
}
