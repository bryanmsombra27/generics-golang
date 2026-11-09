package main

import "testing"

func TestSortAsc(testing *testing.T){
	in:= []int {2,1,3,5,4}
	got:= SortAsc(in)
	want:= []int {1,2,3,4,5,}


	for i:= range want {

		if got[i] !=want[i]{
			testing.Fatalf("got[%d]= %v want=%v",i,got[i],want[i])
		}

	}
}

func TestSortByUsers(testings *testing.T){
	users:= []User{
	{	ID: 1,Name: "Beto"},
	{	ID: 2,Name: "Ana"},
	{	ID: 3,Name: "Fernando"},
	} 

	got:= SortBy(users,func(a,b User) bool {return a.ID < b.ID})
		if got[0].ID != 1 || got[1].ID != 2 || got[2].ID !=3 {
			testings.Fatalf("Orden inesperado %+v",got)
		}


}