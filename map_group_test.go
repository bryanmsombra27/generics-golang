package main

import "testing"

func TestIndexyBy(tesitng *testing.T){

	users:= []User{
		{ID: 1,Name: "Ana"},
		{ID: 2,Name: "Beto"},
		{ID: 3,Name: "Fernando"},
	}

	idx:=IndexBy(users,func(user User) int {return user.ID})

	if idx[1].Name != "Ana" || idx[2].Name !="Beto" || idx[3].Name !="Fernando" {
		tesitng.Fatalf("Index no esperado: %+v",idx)
	}


}

type Order struct {
	ID int
	Customer int
	Total int
}

func TestGroupBy(test *testing.T) {
	orders:= [] Order{
		{ID: 1,Customer: 10,Total: 100},
		{ID: 2,Customer: 10,Total: 200},
		{ID: 3,Customer: 20,Total: 300},
		{ID: 4,Customer: 30,Total: 400},
	}

	got:=GroupBy(orders,func(order Order) int { return order.Customer})
	if len(got[10]) != 2 {
		test.Fatalf("La agrupacion fallo: %+v",got)
	}



}