package generics

import "constraints"

type Predicate[T any] func(T) bool

func And[T any](ps ...Predicate[T]) Predicate[T] {
	return func(v T) bool {
		for _, p := range ps {
			if !p(v) {
				return false
			}
		}
		return true
	}
}

func Or[T any](ps ...Predicate[T]) Predicate[T] {
	return func(v T) bool {
		for _, p := range ps {
			if p(v) {
				return true
			}
		}
		return false
	}
}

func Between[T constraints.Ordered](t1, t2 T) Predicate[T] {
	return func(v T) bool {
		return t1 <= v && v <= t2
	}
}

func LessThan[T constraints.Ordered](t T) Predicate[T] {
	return func(v T) bool {
		return v < t
	}
}

func GreaterThan[T constraints.Ordered](t T) Predicate[T] {
	return func(v T) bool {
		return t < v
	}
}

func Equal[T comparable](t T) Predicate[T] {
	return func(v T) bool {
		return v == t
	}
}

func NotEqual[T comparable](t T) Predicate[T] {
	return func(v T) bool {
		return v != t
	}
}

func StringLength[T ~string](l int) Predicate[T] {
	return func(v T) bool {
		return len(v) == l
	}
}
