package completesearch

var allSubsets [][]int

/*
check this : https://www.youtube.com/watch?v=xFWgZ5DTjFo
*/

func GenereteSubsetVTwo(n int) [][]int {
	for mask := 0; mask < (1 << n); mask++ {
		subset := []int{}
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				subset = append(subset, j)
			}
		}
		allSubsets = append(allSubsets, subset)
	}
	return allSubsets
}
