package completesearch

import (
	"math"
)

func SolveNQueens(n int) [][]string {
	var possibleNQueues [][]int
	var allPossibleNQueues [][]string
	row := 0
	col := 0
	solveNQueensHelper(n, row, col, possibleNQueues, &allPossibleNQueues)
	return allPossibleNQueues
}

func solveNQueensHelper(n int, row int, col int, possibleNQueues [][]int, allPossibleNQueues *[][]string) {
	if len(possibleNQueues) >= n {
		str_array := formatPosition(possibleNQueues, n)
		*allPossibleNQueues = append(*allPossibleNQueues, str_array)
	} else {
		for row := 0; row < n; row++ {
			possibleNQueues = append(possibleNQueues, []int{row, col})
			b := isValidPlacement(possibleNQueues)
			if b {
				solveNQueensHelper(n, row, col+1, possibleNQueues, allPossibleNQueues)
			}
			possibleNQueues = possibleNQueues[:len(possibleNQueues)-1]
		}
	}
}

func isValidPlacement(possibleNQueues [][]int) bool {
	b := true
	lastPlacement := possibleNQueues[len(possibleNQueues)-1]
	for i := 0; i < len(possibleNQueues)-1; i++ {
		placement := possibleNQueues[i]
		diag := math.Abs(float64(placement[0]) - float64(lastPlacement[0]))
		if placement[0] != lastPlacement[0] && placement[1] != lastPlacement[1] && math.Abs(float64(placement[1])-float64(lastPlacement[1])) != diag {
			continue
		} else {
			return false
		}
	}
	return b
}

func formatPosition(possibleNQueues [][]int, n int) []string {
	var possibleNQueuesFormat []string
	for _, v := range possibleNQueues {
		str := ""
		for i := 0; i < n; i++ {
			if i != v[0] {
				str += "."
			} else {
				str += "Q"
			}
		}
		possibleNQueuesFormat = append(possibleNQueuesFormat, str)
	}
	return possibleNQueuesFormat
}
