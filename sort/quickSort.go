package sort

import "algorithms/types"

func QuickSort[T types.Number](items []T, from int, to int) {

	if from < to {
		divider := partition(items, from, to)
		QuickSort(items, from, divider-1)
		QuickSort(items, divider, to)
	}

}

func partition[T types.Number](items []T, from int, to int) int {
	left := from
	right := to
	pivot := items[left]

	for left <= right {
		for items[left] < pivot {
			left++
		}

		for items[right] > pivot {
			right--
		}

		if left <= right {
			items[left], items[right] = items[right], items[left]
			left++
			right--
		}
	}

	return left
}
