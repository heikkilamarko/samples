package generics

import "sort"

func OrderBy[T any](src []T, less func(i, j int) bool) ([]T, error) {
	dst := make([]T, len(src))
	copy(dst, src)
	sort.SliceStable(dst, less)
	return dst, nil
}
