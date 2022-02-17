package generics

import "golang.org/x/exp/constraints"

type Predicate[T any] func(T) (bool, error)

func And[T any](ps ...Predicate[T]) Predicate[T] {
	return func(v T) (bool, error) {
		for _, p := range ps {
			ok, err := p(v)
			if err != nil {
				return false, err
			}
			if !ok {
				return false, nil
			}
		}
		return true, nil
	}
}

func Or[T any](ps ...Predicate[T]) Predicate[T] {
	return func(v T) (bool, error) {
		for _, p := range ps {
			ok, err := p(v)
			if err != nil {
				return false, err
			}
			if ok {
				return true, nil
			}
		}
		return false, nil
	}
}

func Between[T constraints.Ordered](t1, t2 T) Predicate[T] {
	return func(v T) (bool, error) {
		return t1 <= v && v <= t2, nil
	}
}

func LessThan[T constraints.Ordered](t T) Predicate[T] {
	return func(v T) (bool, error) {
		return v < t, nil
	}
}

func GreaterThan[T constraints.Ordered](t T) Predicate[T] {
	return func(v T) (bool, error) {
		return t < v, nil
	}
}

func Equal[T comparable](t T) Predicate[T] {
	return func(v T) (bool, error) {
		return v == t, nil
	}
}

func NotEqual[T comparable](t T) Predicate[T] {
	return func(v T) (bool, error) {
		return v != t, nil
	}
}

func StringLength[T ~string](l int) Predicate[T] {
	return func(v T) (bool, error) {
		return len(v) == l, nil
	}
}
