package main

const CliTemplate = `
package main

import (
	"github.com/urfave/cli"
)

/*
 * configFlags populates and returns a slice with the command line flags
 */
func configFlags() []cli.Flag {
	return []cli.Flag{
{{range .Fields}}
		cli.{{.Meta.CliFlag}}{
			Name:  "{{.Name}}",
			Usage: "{{.Doc}}",
		},{{end}}
	}
}
`
