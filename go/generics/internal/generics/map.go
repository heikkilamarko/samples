package generics

func Map[TSrc, TDst any](src []TSrc, mapper func(TSrc) (TDst, error)) ([]TDst, error) {
	var dst []TDst
	for _, s := range src {
		d, err := mapper(s)
		if err != nil {
			return nil, err
		}
		dst = append(dst, d)
	}
	return dst, nil
}
