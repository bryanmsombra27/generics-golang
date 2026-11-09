package main

import "testing"

func TestContainsAndIndexOf(t *testing.T){
	in:=[]string{"a","b","c"}

	if !Contains(in,"b"){
		t.Fatal("Se esperaba una 'b' ")
	}

	if Contains(in,"z"){
		t.Fatal("No se esperaba 'z'")
	}	
	if IndexOf(in,"c") !=  2 {
		t.Fatal("Se esperaba que la se se encontraera en la posicion 2 del slice")
	}


}

func TestUniquer(t *testing.T){
	in:= []int{1,2,3,4,3,3,4,5,6,6}
	got:= Unique(in)

	want:= []int{1,2,3,4,5,6}


	if len(got) != len(want){
		t.Fatalf("len got=%d want=%d ",len(got),len(want))
	}


for _, i := range want {
	if got[i] != want[i]{
		t.Fatalf("got[%d]= %v  want=%v",i,got[i],want[i])
	} 
}



}