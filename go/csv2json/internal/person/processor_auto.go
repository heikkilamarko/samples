package person

import (
	"csv2json/internal/csv2json"
	"time"

	"github.com/mitchellh/mapstructure"
)

type AutoProcessor struct{}

func NewAutoProcessor() *AutoProcessor {
	return &AutoProcessor{}
}

func (p *AutoProcessor) Process(header, item []string) (interface{}, error) {
	m, err := csv2json.ToMap(header, item)
	if err != nil {
		return nil, err
	}

	c := &mapstructure.DecoderConfig{
		Result:           &Person{},
		WeaklyTypedInput: true,
		DecodeHook:       mapstructure.StringToTimeHookFunc(time.RFC3339),
	}

	d, err := mapstructure.NewDecoder(c)
	if err != nil {
		return nil, err
	}

	if err := d.Decode(m); err != nil {
		return nil, err
	}

	return c.Result, nil
}
