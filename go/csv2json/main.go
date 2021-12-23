package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	processFile("in.csv", "out.json")
}

func processFile(inFilePath, outFilePath string) {
	inFile, err := os.Open(inFilePath)
	checkErr(err)

	defer inFile.Close()

	r := csv.NewReader(inFile)
	r.ReuseRecord = true
	r.LazyQuotes = true

	outFile, err := os.Create(outFilePath)
	checkErr(err)

	defer outFile.Close()

	w := bufio.NewWriter(outFile)

	header, err := readHeader(r)
	checkErr(err)

	for {
		inItem, err := r.Read()

		if err == io.EOF {
			break
		}

		checkErr(err)

		outItem := processItem(header, inItem)

		bytes, err := json.Marshal(outItem)
		checkErr(err)

		w.Write(bytes)
		w.WriteString("\n")
	}

	w.Flush()
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

func processItem(header, item []string) interface{} {
	m := map[string]interface{}{}

	for i, v := range item {
		m[header[i]] = v
	}

	return m
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
