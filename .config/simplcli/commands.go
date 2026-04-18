package main

import "fmt"

type commandsGroup []command
type commandFunc func(...string) int8

type command struct {
	Name  string
	Descr string
	Func  commandFunc
}

func (cg commandsGroup) Run(args ...string) int8 {
	m := make(map[string]*command)
	for _, c := range cg {
		m[c.Name] = &c
	}
	if len(args) >= 1 {
		n := args[0]
		if c := m[n]; c != nil {
			return c.Func(args[1:]...)
		} else {
			fmt.Printf("command %s not found\n", n)
		}
	}
	for _, c := range cg {
		fmt.Printf("%s - %s\n", c.Name, c.Descr)
	}
	return 1
}

func runList(cs ...command) int8 {
	for _, c := range cs {
		if r := c.Func(); r != 0 {
			return r
		}
	}
	return 0
}
