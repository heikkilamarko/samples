package main

import (
	"csv2json/internal/csv2json"
	"csv2json/internal/person"
	"flag"
	"fmt"
	"os"
)

func main() {
	i, o, p := parseFlags()

	var ip csv2json.ItemProcessor

	if p == "auto" {
		ip = person.NewAutoProcessor()
	} else {
		ip = person.NewProcessor()
	}

	processor := csv2json.NewProcessor(ip)

	fmt.Printf("Processing %s -> %s (processor: '%s') ...\n", i, o, p)

	if err := processor.ProcessFile(i, o); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("...done")
}

func parseFlags() (string, string, string) {
	var i, o, p string

	fs := flag.NewFlagSet("csv2json", flag.ExitOnError)

	fs.StringVar(&i, "i", "", "input file (required)")
	fs.StringVar(&o, "o", "", "output file")
	fs.StringVar(&p, "p", "default", "person processor ('default' or 'auto')")

	fs.Parse(os.Args[1:])

	if i == "" {
		fs.Usage()
		os.Exit(2)
	}

	if o == "" {
		o = i + ".json"
	}

	if p != "auto" {
		p = "default"
	}

	return i, o, p
}
