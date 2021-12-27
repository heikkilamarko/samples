package csv2json

import "errors"

func ToMap(header, item []string) (map[string]string, error) {
	if len(header) < len(item) {
		return nil, errors.New("header length is less than item length")
	}

	m := map[string]string{}

	for i, v := range item {
		m[header[i]] = v
	}

	return m, nil
}
