package main

import "testing"

func TestStockMarket(t *testing.T) {
	set := []int{1,2,3,4}
	buyIndex,sellIndex, err := StockMarket(set)

	if buyIndex != 0 && sellIndex != 3 && err != nil {
		t.Fail()
	}

	set = []int{1,2,3,4,5}
	buyIndex,sellIndex,err = StockMarket(set)

	if buyIndex != 0 && sellIndex != 4 && err != nil {
		t.Fail()
	}

	set = []int{10,2,3,4,6}
	buyIndex,sellIndex,err = StockMarket(set)

	if buyIndex != 1 && sellIndex != 4 && err != nil {
		t.Fail()
	}
}