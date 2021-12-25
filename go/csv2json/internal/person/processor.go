package person

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Process(header, item []string) (interface{}, error) {
	m := map[string]string{}

	for i, v := range item {
		m[header[i]] = v
	}

	person := &Person{}

	if err := p.decode(m, person); err != nil {
		return nil, err
	}

	return person, nil
}

func (p *Processor) decode(from map[string]string, to *Person) error {
	c := &mapstructure.DecoderConfig{
		Result:           to,
		WeaklyTypedInput: true,
		DecodeHook:       mapstructure.StringToTimeHookFunc(time.RFC3339),
	}

	d, err := mapstructure.NewDecoder(c)
	if err != nil {
		return err
	}

	return d.Decode(from)
}
