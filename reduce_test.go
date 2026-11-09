package main

import "testing"

func TestReduceSum(t *testing.T) {
	nums:=[]int{ 1,2,3,4}

	sum:= Reduce(nums,0,func(acc, num int)int {
		return acc +num
	})

	if sum != 10 {
		t.Fatalf("La funcion reduce fallo want=[10] got=[%d]",sum)
	}
}

func TestPipelineOrder(test *testing.T){
		orders:= [] Order{
		{ID: 1,Customer: 10,Total: 100},
		{ID: 2,Customer: 10,Total: 200},
		{ID: 3,Customer: 20,Total: 300},
		{ID: 4,Customer: 30,Total: 400},
		{ID: 5,Customer: 10,Total: 500},
	}

	customer10:= Filter(orders,func(order Order) bool {
		return order.Customer == 10
	})

	totals:= Map(customer10,func(order Order) int {return order.Total})

	sum:= Reduce(totals,0,func(acc,number int) int {return acc + number})


	if sum != 800 {
		test.Fatalf("Queria 300  obtuve: %d",sum)
	}

}