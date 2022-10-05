package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	var test = []int{4, 7, 5, 3, 6, 8, 9, 61, 4, 7, 5, 3, 5, 7}
	sort.QuickSort(test, 0, len(test)-1)

	fmt.Println(test)
}
