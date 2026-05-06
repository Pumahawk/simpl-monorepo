package cmd

import (
	"errors"
	"flag"
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

type CommandGroup struct {
	Name        string
	Commands    []CommandW
	FlagFunc    func(*flag.FlagSet)
	FlagValFunc func() error
}

func (c *CommandGroup) CRun(args []string) (any, error) {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	if c.FlagFunc != nil {
		c.FlagFunc(fs)
	}

	fs.Parse(args)
	args = fs.Args()

	if c.FlagValFunc() != nil {
		if err := c.FlagValFunc(); err != nil {
			return nil, err
		}
	}

	if len(args) > 0 {
		for _, c := range c.Commands {
			if c.CName() == args[0] {
				return c.CRun(args[1:])
			}
		}
		return nil, fmt.Errorf("command %q not found\n", args[0])
	} else {
		fmt.Printf("commands:\n")
		for _, c := range c.Commands {
			fmt.Printf("%s\n", c.CName())
		}
		return nil, MissingCommand
	}
}

func (c *CommandGroup) CName() string {
	return ""
}
