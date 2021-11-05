package main

import (
	"fmt"
)

func main() {
	array := []int{1, 4, 9, 6, 8, 2, 2}
	quickSorting(array, 0, len(array))
	fmt.Println(array)

}

func partition(array []int, start int, end int) int {
	pivot := array[start]
	k := start + 1
	for i := start + 1; i < end; i++ {
		if array[i] < pivot {
			x := array[i]
			array[i] = array[k]
			array[k] = x
			k++
		}
		array[start] = array[k-1]
		array[k-1] = pivot
	}

	return k - 1
}

func quickSorting(array []int, start int, end int) {
	if start == end {
		return
	}
	index := partition(array, start, end)
	quickSorting(array, start, index)
	quickSorting(array, index+1, end)
}
