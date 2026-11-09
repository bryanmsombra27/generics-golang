package main

import (
	"fmt"
	"testing"
)

func TestPartition(test *testing.T) {

	numbers:= []int{1,2,3,4,5,6,7,}

	pair,even := Partition(numbers,func(number int) bool {
		return number %2 == 0
	})


	if len(pair) != 3 {
	test.Fatalf("total de numeros pares 3, recibidos:[%d]",len(pair))
}
if len(even) != 4 {
	
	test.Fatalf("total de numeros impares 4, recibidos:[%d]",len(even))
	}
	fmt.Printf("Numeros pares: [%v]",pair)
	fmt.Printf("Numeros impares: [%v]",even)
	

}