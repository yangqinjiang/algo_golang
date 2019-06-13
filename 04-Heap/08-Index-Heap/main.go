package main

import (
	"algo_golang/02-Sorting-Basic/core"
	core2 "algo_golang/04-Heap/core"
)

func main() {
	n := 1000000

	st := new(core.SortTestHelper)
	arr1 := st.GenerateRandomArray(n, 0, n)
	arr2 := st.CopyIntArray(arr1, n)
	st.TestSort("Heap Sort Using Index-Max-Heap",core2.HeapSortUsingIndexMaxHeap,arr1,n)
	st.TestSort("Heap Sort 4", core2.HeapSort4, arr2, n)
}
