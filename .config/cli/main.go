package main

import (
	"fmt"
	"os"
	"text/template"
)

var commands = []Command{
	KeycloakRealmAutoconfigure,
}

func main() {
	if len(os.Args) > 1 {
		cname := os.Args[1]
		cargs := os.Args[2:]
		for _, command := range commands {
			if command.Name == cname {
				command.Run(cargs)
				return
			}
		}
		fmt.Printf("Unknow subcommand: %s\n\n", cname)
	}
	PrintHelp()
}

func PrintHelp() {
	t := template.Must(template.New("help-message").
		Parse(`subcommand:
{{- range .Commands }}
- {{ .Name }}: {{ .Description }}
{{- end }}
- help: Print this message
`))
	v := struct {
		Commands []Command
	}{
		Commands: commands,
	}
	t.Execute(os.Stdout, v)
}
