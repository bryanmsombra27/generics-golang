package main

func Contains[T comparable](in []T, target T) bool {

	for _, value := range in {
		if value == target {
			return true
		}
	}
	return false
}

func Unique[T comparable](in []T) []T {

	seen := make(map[T]struct{}, len(in))
	out := make([]T, 0, len(in))

	for _, value := range in {
		if _, ok := seen[value]; ok {
			continue
		}

		seen[value] = struct{}{}
		out = append(out, value)
	}

	return out

}

func IndexOf[T comparable](slice []T, search T) int {

	for index, value := range slice {

		if search == value {
			return index
		}
	}
	return -1
}
