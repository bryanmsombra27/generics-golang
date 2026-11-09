package main

import (
	"cmp"
	"slices"
)

// tipo de dato ordenable
func SortAsc[T cmp.Ordered](in []T) []T {
	out:=slices.Clone(in)

	slices.Sort(out)
	return out
	

}


// ordenar cualquier tipo de dato
func SortBy[T any](in []T, less func(a,b T) bool ) []T {
	out:= slices.Clone(in)

	slices.SortFunc(out,func(a,b T) int {
		if less(a,b) {
			return -1
		}
		if less(b,a){
			return 1
		}
		return 0
	})

	return out	
}
