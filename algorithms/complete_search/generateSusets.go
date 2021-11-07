package completesearch

import "fmt"

var subset []int

func GenereteSubset(c []int, k int) {
	if k == len(c) {
		fmt.Printf("subset : %v\n", subset)
	} else {
		//fmt.Println(k, &subset)
		GenereteSubset(c, k+1)
		subset = append(subset, c[k])
		GenereteSubset(c, k+1)
		subset = subset[:len(subset)-1]

	}

}

var allsubsets [][]int

func GenereteSubsets(k int, n int) {
	if k == n {
		subsetCopy := make([]int, n)
		nbytes := copy(subsetCopy, subset)
		allsubsets = append(allsubsets, subsetCopy[:nbytes])
	} else {
		subset = append(subset, k)
		GenereteSubsets(k+1, n)
		subset = subset[:len(subset)-1]
		GenereteSubsets(k+1, n)
	}
}

func GenereteSubsetsPrinter(k int, n int) [][]int {
	subset = []int{}
	GenereteSubsets(k, n)
	return allsubsets
}
