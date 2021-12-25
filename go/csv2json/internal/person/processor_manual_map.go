package person

import (
	"strconv"
	"time"
)

type ManualMapProcessor struct{}

func NewManualMapProcessor() *ManualMapProcessor {
	return &ManualMapProcessor{}
}

func (p *ManualMapProcessor) Process(header, item []string) (interface{}, error) {
	m := map[string]string{}

	for i, v := range item {
		m[header[i]] = v
	}

	var err error

	person := &Person{
		Name: m["name"],
	}

	person.Age, err = strconv.Atoi(m["age"])
	if err != nil {
		return nil, err
	}

	person.Height, err = strconv.ParseFloat(m["height"], 64)
	if err != nil {
		return nil, err
	}

	person.IsActive, err = strconv.ParseBool(m["is_active"])
	if err != nil {
		return nil, err
	}

	person.CreatedAt, err = time.Parse(time.RFC3339, m["created_at"])
	if err != nil {
		return nil, err
	}

	return person, nil
}
