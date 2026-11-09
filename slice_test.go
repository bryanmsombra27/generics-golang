package main

import "testing"


func TestFilterInt(test *testing.T){
	in:=[]int{1,2,3,4,5}

	got:=Filter(in,func(number int)bool {
		return number %2 ==0
	})

	want:= []int{2,4}

	if(len(got) != len(want)){
		// fatalf hace que el test falle 
		test.Fatalf("Len got %d want %d",len(got),len(want))
	}

	for i:=range want {
		if(got[i] != want[i]){
			test.Fatalf("got [%d]=%v want=%v",i,got[i],want[i])
		}
	}


}

type User struct {
	ID int
	Name string

}


func TestMapToNames(test *testing.T){
	users:=[]User{
		{ID: 1 ,Name: "Ana" },
		{ID: 2 ,Name: "Beto" },
		{ID: 3 ,Name: "Fernando" },
	}

	got:=Map(users,func(user User) string {
		return user.Name

	})

	want:= []string{
		"Ana",
		"Beto",
		"Fernando",
	}

	if(len(got)!= len(want)){
		test.Fatalf("Los valores no coinciden")
	}	
	

}	