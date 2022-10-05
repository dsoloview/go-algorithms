package sort

import (
	"algorithms/types"
)

func MergeSort[T types.Number](items []T) []T {
	if len(items) == 1 {
		return items
	}

	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])

	return merge(first, second)
}

func merge[T types.Number](first []T, second []T) []T {
	var result []T
	var i, j int
	for i < len(first) && j < len(second) {
		if first[i] > second[j] {
			result = append(result, second[j])
			j++
		} else {
			result = append(result, first[i])
			i++
		}
	}

	for ; i < len(first); i++ {
		result = append(result, first[i])
	}

	for ; j < len(second); j++ {
		result = append(result, second[j])
	}

	return result
}
