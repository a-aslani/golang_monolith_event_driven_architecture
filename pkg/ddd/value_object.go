package ddd

type ValueObject[T any] interface {
	Value() T
}
