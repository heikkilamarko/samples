package generics

func IsEqual[T comparable](v1, v2 T) bool {
	return v1 == v2
}

func IsNotEqual[T comparable](v1, v2 T) bool {
	return v1 != v2
}
