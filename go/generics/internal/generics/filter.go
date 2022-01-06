package generics

func Filter[T any](src []T, p Predicate[T]) ([]T, error) {
	var dst []T
	for _, s := range src {
		ok, err := p(s)
		if err != nil {
			return nil, err
		}
		if ok {
			dst = append(dst, s)
		}
	}
	return dst, nil
}
