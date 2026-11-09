package main

func Sum[T Number](slice []T) T {

	var acc T

	for _, number := range slice {
		acc += number
	}

	return acc
}