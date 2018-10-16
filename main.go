package main

import "fmt"

func main() {

	//unsortedSet1 := []int{0,80,13,4,54,2,6,1,4,3,7,23,8,65,33,23,9,3,1,76}
	unsortedSet1 := []int{80,76,65,54,33,23,23,13,9,8,7,6,4,4,3,3,2,1,1,0}
	fmt.Printf("UNSORTED SET 1:       %v\n\n",unsortedSet1)

	fmt.Printf("INSERTION SORT SET 1: %v\n",insertionSort(unsortedSet1))
	//fmt.Printf("SELECTION SORT SET 1: %v\n",selectionSort(unsortedSet1))
	fmt.Println("SORTED SET 1:         [0 1 1 2 3 3 4 4 6 7 8 9 13 23 23 33 54 65 76 80]")
	//fmt.Printf("BUBBLE SORT SET 1:    %v\n",bubbleSort(unsortedSet1,2))

}

func insertionSort(set []int) ([]int) {
	for i := 0; i < len(set); i++ {
		j := i
		for j > 0 && set[j-1] > set[j] {
			swap(&set,j-1,j)
			j--
		}
	}
	return set
}

func selectionSort (set []int) ([]int) {
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

func bubbleSort(set []int, iterations int) ([]int){
	for i := 0 ; i < iterations; i++ {
		for n := 0; n < len(set)-1; n++ {
			if set[n] > set[n+1] {
				swap(&set,n,n+1)
			}
		}
	}
	return set
}

func swap(set *[]int, index1 int, index2 int) {
	temp := (*set)[index1]
	(*set)[index1] = (*set)[index2]
	(*set)[index2] = temp
}
