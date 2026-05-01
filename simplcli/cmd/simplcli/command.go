package main

import (
	"errors"
	"fmt"
)

var MissingCommand = errors.New("missing command")

type CommandFunc[T any] func(c *Command[T], args []string) (T, error)

// Usefull to groups differents commands
type CommandW interface {
	CName() string
	CRun(args []string) (any, error)
}

type Command[T any] struct {
	Name string
	Run  CommandFunc[T]
}

// Implementation CommandW
func (c *Command[T]) CName() string {
	return c.Name
}

// Implementation CommandW
func (c *Command[T]) CRun(args []string) (any, error) {
	return c.Run(c, args)
}

func CommandGroup(name string, commands ...CommandW) CommandW {
	return &Command[any]{
		Name: name,
		Run: func(c *Command[any], args []string) (any, error) {
			if len(args) > 0 {
				for _, c := range commands {
					if c.CName() == args[0] {
						return c.CRun(args[1:])
					}
				}
				return nil, fmt.Errorf("command %q not found\n", args[0])
			} else {
				fmt.Printf("commands:\n")
				for _, c := range commands {
					fmt.Printf("%s\n", c.CName())
				}
				return nil, MissingCommand
			}
		},
	}
}
