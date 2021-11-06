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
