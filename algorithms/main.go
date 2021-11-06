package main

import (
	"fmt"

	bitmask "github.com/souben/algorithms/bit_mask"
	completesearch "github.com/souben/algorithms/complete_search"
	quicksorting "github.com/souben/algorithms/quick_sorting"
)

func main() {
	array := []int{1, 4, 9, 6, 8, 2, 2}
	quicksorting.QuickSorting(array, 0, len(array))
	fmt.Println(array)

	array = []int{1, 4, 9, 6, 8, 2, 3}
	quicksorting.QuickSorting_(array, 0, len(array))
	fmt.Println(array)

	c := []int{1, 2, 3}
	k := 0
	completesearch.GenereteSubset(c, k)

	bitmask.CombineAndOr()
}
