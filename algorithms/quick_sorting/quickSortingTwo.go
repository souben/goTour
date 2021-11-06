package quicksorting

func partition_(array []int, start int, end int) int {
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
	array[end-1] = array[k+1]
	array[k+1] = pivot
	return k + 1
}

func QuickSorting_(array []int, start int, end int) {
	if start >= end-1 {
		return
	}
	index := partition_(array, start, end)
	QuickSorting_(array, start, index)
	QuickSorting_(array, index, end)
}
