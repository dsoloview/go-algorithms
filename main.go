package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	var test = []int{4, 7, 5, 3, 6, 8, 9, 61, 4, 7, 5, 3, 5, 7}
	result := sort.MergeSort(test)

	fmt.Println(test)
	fmt.Println(result)
}
