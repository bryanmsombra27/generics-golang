package main

import "testing"

func TestSum(test *testing.T){

	if Sum([]int{1,2,3}) != 6{
		test.Fatalf("Se esperaba un 6")
	}

	if Sum([]float64{1.5,2.5}) != 4.0 {
		test.Fatalf("se esperaba un 4.0")
	}



}
