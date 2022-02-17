package generics

import "golang.org/x/exp/slices"

func OrderBy[T any](src []T, less func(a, b T) bool) ([]T, error) {
	dst := make([]T, len(src))
	copy(dst, src)
	slices.SortStableFunc(dst, less)
	return dst, nil
}
