package person

import (
	"strconv"
	"time"
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

	var err error

	person := &Person{Name: m["name"]}

	if person.Id, err = strconv.Atoi(m["id"]); err != nil {
		return nil, err
	}

	if person.Age, err = strconv.Atoi(m["age"]); err != nil {
		return nil, err
	}

	if person.Height, err = strconv.ParseFloat(m["height"], 64); err != nil {
		return nil, err
	}

	if person.IsActive, err = strconv.ParseBool(m["is_active"]); err != nil {
		return nil, err
	}

	if person.CreatedAt, err = time.Parse(time.RFC3339, m["created_at"]); err != nil {
		return nil, err
	}

	return person, nil
}
