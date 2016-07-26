package main

import (
	"flag"
	"fmt"
	"github.com/aurelien-rainone/gotypes"
	"io/ioutil"
	"log"
	"os"
)

var (
	typename = flag.String("type", "", "type name to inspect; must be set")
	outname  = flag.String("out", "", "name of the generated file; if empty the eventually not correct output is printed on standard output")
	filename string
)

func showUsage() {
	fmt.Println("gencli - Automatically generate command-line interface for Surviveler")
	fmt.Println()
	fmt.Println("usage:")
	fmt.Println("  gencli -type string -out string inputfile")
	fmt.Println()
	fmt.Println("options:")
	fmt.Println("  inputfile")
	fmt.Println("        go input file; must not be set if called from `go generate`")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gencli: ")
	flag.Usage = showUsage
	flag.Parse()

	if *typename == "" {
		flag.Usage()
		os.Exit(1)
	}

	file := ""
	if flag.NArg() > 0 {
		file = flag.Arg(0)
	} else {
		file = os.Getenv("GOFILE")
	}

	tdef, err := gotypes.Inspect(*typename, file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range tdef.Fields {
		// add 'CliFlag' to the 'Meta' container
		f.Meta["CliFlag"] = cliFlagFromType(f.Type)
	}

	buf := gotypes.Generate(CliTemplate, tdef, *outname == "")
	if *outname != "" {
		// generate output filename
		err = ioutil.WriteFile(*outname, buf, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}
}

func cliFlagFromType(t string) string {
	switch t {
	case "string":
		return "StringFlag"
	case "int":
		return "IntFlag"
	default:
		log.Fatalf("unsupported struct field type: %s\n", t)
		return ""
	}
}
