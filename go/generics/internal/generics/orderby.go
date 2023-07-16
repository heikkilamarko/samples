package generics

import "slices"

func OrderBy[T any](src []T, cmp func(a, b T) int) ([]T, error) {
	dst := make([]T, len(src))
	copy(dst, src)
	slices.SortStableFunc(dst, cmp)
	return dst, nil
}
