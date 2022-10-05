package sort

import "algorithms/types"

func MergeSort[T types.Number](items []T) []T {
	if len(items) == 1 {
		return items
	}

	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])
	return merge(first, second)

}

func merge[T types.Number](first []T, seconds []T) []T {
	var result []T
	var i, j int
	for i < len(first) && j < len(seconds) {
		if first[i] < seconds[j] {
			result = append(result, first[i])
			i++
		} else {
			result = append(result, seconds[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		result = append(result, first[i])
	}
	for ; j < len(seconds); j++ {
		result = append(result, seconds[j])
	}

	return result
}
