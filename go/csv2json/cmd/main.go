package main

import (
	"csv2json/internal/csv2json"
	"csv2json/internal/person"
	"flag"
	"fmt"
	"os"
)

func main() {
	i, o := parseFlags()

	p := csv2json.NewProcessor(person.NewProcessor())

	if err := p.ProcessFile(i, o); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseFlags() (string, string) {
	var i, o string

	fs := flag.NewFlagSet("csv2json", flag.ExitOnError)

	fs.StringVar(&i, "i", "", "input file (required)")
	fs.StringVar(&o, "o", "", "output file")

	fs.Parse(os.Args[1:])

	if i == "" {
		fs.Usage()
		os.Exit(2)
	}

	if o == "" {
		o = i + ".json"
	}

	return i, o
}
