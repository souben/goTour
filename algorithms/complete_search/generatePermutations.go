package completesearch

import "fmt"

var perm []int
var allpermutations [][]int
var chosen []bool = make([]bool, 4)

func GeneratePermutations(n int) {
	if len(perm) == n {

		fmt.Println(perm)
	}
	for i := 0; i < n; i++ {
		if chosen[i] {
			continue
		}
		chosen[i] = true
		perm = append(perm, i)
		//fmt.Println(perm, i, "*")
		GeneratePermutations(n)
		chosen[i] = false
		perm = perm[:len(perm)-1]
		//fmt.Println(perm, i, "*")
	}
}

func GeneratePermutationsIterative(n int) {
	allpermutations = append(allpermutations, []int{0})
	for i := 1; i < n; i++ {
		a := i
		var results [][]int
		for len(allpermutations) > 0 {
			b := allpermutations[0]
			allpermutations = allpermutations[1:]
			for k := 0; k < len(b)+1; k++ {
				l := make([]int, len(b))
				var g []int
				copy(l, b)
				g = append(g, l[:k]...)
				g = append(g, a)
				g = append(g, l[k:]...)
				results = append(results, g)
			}
		}
		allpermutations = append(allpermutations, results...)
	}
	for _, v := range allpermutations {
		fmt.Println(v)
	}
}
