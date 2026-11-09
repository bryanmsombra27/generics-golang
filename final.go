package main

func Partition[T any](items []T, fn func(number T) bool) ([]T, []T) {
	//creacion 2 slice vacios
	truthly := make([]T, 0)
	falsy := make([]T, 0)

	for _, value := range items {
		if fn(value) {
			truthly = append(truthly, value)
		} else {
			falsy = append(falsy, value)
		}

	}

	return truthly, falsy
}