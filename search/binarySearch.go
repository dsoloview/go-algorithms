package search

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* https://stepik.org/lesson/13246/step/4?unit=3431 */
func MainBinarySearch() {

	var resultStringSlice []string

	arr, err := readInput()
	if err != nil {
		panic(err)
	}

	search, err := readInput()
	if err != nil {
		panic(err)
	}

	for _, value := range search[1:] {
		result := binarySearch(arr, value)
		resultStringSlice = append(resultStringSlice, strconv.Itoa(result))
	}

	fmt.Println(strings.Join(resultStringSlice, " "))
}

func readInput() ([]int, error) {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()

	inputSlice := readNumbers(input)

	return inputSlice, nil

}

func readNumbers(s string) []int {
	var numbers []int
	for _, value := range strings.Fields(s) {
		i, err := strconv.Atoi(value)
		if err == nil {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

func binarySearch(arr []int, k int) int {
	first := 1
	last := len(arr) - 1

	for last >= first {
		mid := (first + last) / 2
		if arr[mid] == k {
			return mid
		}
		if k > arr[mid] {
			first = mid + 1
		}

		if k < arr[mid] {
			last = mid - 1
		}
	}

	return -1
}
