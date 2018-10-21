package main

import (
	"fmt"
	"log"
	"math"
)

func main() {

	unsortedSet1 := []int{80,76,65,54,33,23,23,13,9,8,7,6,4,4,3,3,2,1,1,0}

	fmt.Printf("UNSORTED SET 1:       %v\n\n",unsortedSet1)
	fmt.Println("SORTED SET 1:         [0 1 1 2 3 3 4 4 6 7 8 9 13 23 23 33 54 65 76 80]")
	fmt.Printf("INSERTION SORT SET 1: %v\n",insertionSort(unsortedSet1))
	fmt.Printf("SELECTION SORT SET 1: %v\n",selectionSort(unsortedSet1))
	fmt.Printf("BUBBLE SORT SET 1:    %v\n",bubbleSort(unsortedSet1,len(unsortedSet1)))
	fmt.Printf("MERGE  SORT SET 1:    %v\n",mergeSort(unsortedSet1))

	fmt.Print("\nStock Market Problem:\n")

	stockSet := []int{100,2,3,4,5,6,1,90}

	buyIndex, sellIndex, err := StockMarket(stockSet)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCalculated...\nBuyIndex: %v, SellIndex: %v\nBuyValue: %v, SellValue: %v\nTotal Profit: %v\n",buyIndex,sellIndex,stockSet[buyIndex],stockSet[sellIndex],stockSet[sellIndex]-stockSet[buyIndex])
}

func insertionSort(rawSet []int) ([]int) {
	set := make([]int, len(rawSet)) //required because of the way slices work. definitely would not do this on production hahah
	copy(set, rawSet)

	for i := 0; i < len(set); i++ {
		j := i
		for j > 0 && set[j-1] > set[j] {
			swap(&set,j-1,j)
			j--
		}
	}
	return set
}

func selectionSort (rawSet []int) ([]int) {
	set := make([]int, len(rawSet))//required because of the way slices work. definitely would not do this on production hahah
	copy(set, rawSet)
	for i := 0; i < len(set); i++ {
		minValue := set[i]
		minIndex := i
		for j := i; j < len(set); j++ {
			if set[j] < minValue {
				minValue = set[j]
				minIndex = j
			}
		}
		swap(&set,i,minIndex)
	}
	return set
}

func bubbleSort(rawSet []int, iterations int) ([]int){
	set := make([]int, len(rawSet))//required because of the way slices work. definitely would not do this on production hahah
	copy(set, rawSet)

	for i := 0 ; i < iterations; i++ {
		for n := 0; n < len(set)-1; n++ {
			if set[n] > set[n+1] {
				swap(&set,n,n+1)
			}
		}
	}
	return set
}

func mergeSort(rawSet []int) ([]int) {
	set := make([]int, len(rawSet))//required because of the way slices work. definitely would not do this on production hahah
	copy(set, rawSet)

	if len(set) == 1 {
		 return set
	}

	leftSide := mergeSort(set[:len(set)/2])
	rightSide := mergeSort(set[len(set)/2:])

	newSet := merge(leftSide,rightSide)

	return newSet
}

func merge(A []int, B []int) ([]int) {
	newSet := make([]int, len(A) + len(B))
	A = append(A,math.MaxInt64)
	B = append(B,math.MaxInt64)
	var i,k int
	for curIndex := 0; curIndex < (len(A) + len(B)) - 2; curIndex++ { //subtract 2 because we add 2 elements for simplicity sake
		if A[i] < B[k] {
			newSet[curIndex] = A[i]
			i++
		} else {
			newSet[curIndex] = B[k]
			k++
		}
	}
	return newSet
}

func swap(set *[]int, index1 int, index2 int) {
	temp := (*set)[index1]
	(*set)[index1] = (*set)[index2]
	(*set)[index2] = temp
}
