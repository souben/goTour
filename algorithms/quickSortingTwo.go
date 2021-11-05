package main

import (
	"fmt"
)

func main() {
	array := []int{1, 4, 9, 6, 8, 2, 3}
	quickSorting(array, 0, len(array))
	fmt.Println(array)

}

func partition(array []int, start int, end int) int {
	pivot := array[end-1]
	k := end - 2
	for i := end - 2; i >= start; i-- {
		if array[i] > pivot {
			x := array[i]
			array[i] = array[k]
			array[k] = x
			k--
		}
	}
	fmt.Println(array, pivot)
	array[end-1] = array[k+1]
	array[k+1] = pivot
	fmt.Println(array)
	return k + 1
}

func quickSorting(array []int, start int, end int) {
	if start >= end-1 {
		return
	}
	index := partition(array, start, end)
	fmt.Println(index, start, end)
	quickSorting(array, start, index)
	quickSorting(array, index, end)
}
