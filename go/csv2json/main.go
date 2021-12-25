package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	i, o := parseFlags()

	if err := processFile(i, o); err != nil {
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

func processFile(inFilePath, outFilePath string) error {
	inFile, err := os.Open(inFilePath)
	if err != nil {
		return err
	}

	defer inFile.Close()

	r := csv.NewReader(inFile)
	r.ReuseRecord = true
	r.LazyQuotes = true

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}

	defer outFile.Close()

	w := bufio.NewWriter(outFile)

	defer w.Flush()

	header, err := readHeader(r)
	if err != nil {
		return err
	}

	for {
		inItem, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		outItem, err := processItem(header, inItem)
		if err != nil {
			return err
		}

		bytes, err := json.Marshal(outItem)
		if err != nil {
			return err
		}

		w.Write(bytes)
		w.WriteString("\n")
	}

	return nil
}

func readHeader(r *csv.Reader) ([]string, error) {
	h, err := r.Read()
	if err != nil {
		return nil, err
	}

	header := make([]string, len(h))
	copy(header, h)

	return header, nil
}

func processItem(header, item []string) (interface{}, error) {
	m := map[string]interface{}{}

	for i, v := range item {
		m[header[i]] = v
	}

	return m, nil
}
