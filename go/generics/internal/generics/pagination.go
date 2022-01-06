package generics

import "errors"

func Page[T any](src []T, offset, limit int) ([]T, error) {
	if offset < 0 {
		return nil, errors.New("invalid offset")
	}

	if limit < 1 {
		return nil, errors.New("invalid limit")
	}

	var dst []T

	if len(src) <= offset {
		return dst, nil
	}

	if limit < len(src)-offset {
		dst = src[offset : offset+limit]
	} else {
		dst = src[offset:]
	}

	return dst, nil
}
