package sort

import "algorithms/types"

func BubbleSort[T types.Number](slice []T) {
	sliceLen := len(slice)
	for i := 0; i < sliceLen; i++ {
		for j := i + 1; j < sliceLen; j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
