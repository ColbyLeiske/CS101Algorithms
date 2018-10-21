package main

import (
	"errors"
	"math"
)

//O(nlogn)
//To get to O(logn) we'd have to get min and max values index from our recursive calls.
func StockMarket(set []int) (buyIndex int, sellIndex int, err error) {
	if len(set) == 2 {
		return 0,1,nil
	}
	if len(set) == 1 {
		return 0,0, errors.New("length of input is 1")
	}

	buyIndexLeft, sellIndexLeft, errLeft := StockMarket(set[:len(set)/2])
	buyIndexRight, sellIndexRight, errRight := StockMarket(set[(len(set)/2):])
	buyIndexRight += len(set)/2
	sellIndexRight += len(set)/2

	var buyIndexLeftFull int
	for index := range set[:len(set)/2] {
		if set[index] < set[buyIndexLeftFull] {
			buyIndexLeftFull = index
		}
	}

	sellIndexRightFull := len(set)/2 //it all needs the len(set)/2 offset because of translation from local set to global set
	for index := range set[len(set)/2:] {
		if set[index + len(set)/2] > set[sellIndexRightFull] {
			sellIndexRightFull = index + len(set)/2
		}
	}

	leftProfit := set[sellIndexLeft] - set[buyIndexLeft]
	rightProfit := set[sellIndexRight] - set[buyIndexRight]
	fullProfit := set[sellIndexRightFull] - set[buyIndexLeftFull]

	if errLeft != nil {
		leftProfit = math.MinInt64 //certainly not the best method of handling this
	}
	if errRight != nil {
		rightProfit = math.MinInt64
	}

	if leftProfit > rightProfit && leftProfit > fullProfit {
		return buyIndexLeft,sellIndexLeft,nil
	}
	if rightProfit > leftProfit && rightProfit > fullProfit {
		return buyIndexRight,sellIndexRight,nil
	}

	return buyIndexLeftFull,sellIndexRightFull,nil //Assumes the left AND right aren't larger
}
