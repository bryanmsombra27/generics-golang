package main

func Reduce[T any, R any](items []T, init R, fn func(accumulate R, item T) R) R {
	accumulate := init

	for _, value := range items {
		accumulate = fn(accumulate, value)
	}
	return accumulate
}