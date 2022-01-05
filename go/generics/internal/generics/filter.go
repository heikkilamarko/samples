package generics

func Filter[T any](vs []T, p Predicate[T]) []T {
	var res []T
	for _, v := range vs {
		if p(v) {
			res = append(res, v)
		}
	}
	return res
}
