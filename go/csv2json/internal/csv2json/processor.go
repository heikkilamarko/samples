package csv2json

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
)

type ItemProcessor interface {
	Process(header, item []string) (interface{}, error)
}

type Processor struct {
	ip ItemProcessor
}

func NewProcessor(ip ItemProcessor) *Processor {
	return &Processor{ip}
}

func (p *Processor) ProcessFile(inFilePath, outFilePath string) error {
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

	header, err := p.readHeader(r)
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

		outItem, err := p.ip.Process(header, inItem)
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

func (p *Processor) readHeader(r *csv.Reader) ([]string, error) {
	h, err := r.Read()
	if err != nil {
		return nil, err
	}

	header := make([]string, len(h))
	copy(header, h)

	return header, nil
}
