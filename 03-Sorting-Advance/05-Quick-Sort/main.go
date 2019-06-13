package main

import (
	"algo_golang/02-Sorting-Basic/core"
	core2 "algo_golang/03-Sorting-Advance/core"
	"fmt"
)

func main() {

	fmt.Println("running Quick Sort")

	// 比较Merge Sort和Quick Sort两种排序算法的性能效率
	// 两种排序算法虽然都是O(nlogn)级别的, 但是Quick Sort算法有常数级的优势
	// Quick Sort要比Merge Sort快, 即使我们对Merge Sort进行了优化

	n := 1000000
	ms := new(core2.MergeSort)
	qs := new(core2.QuickSort)
	// 测试1 一般性测试
	fmt.Printf("Test for random array,size = %d , random range [0,%d]\n", n, n)

	st := new(core.SortTestHelper)
	arr1 := st.GenerateRandomArray(n, 0, n)
	arr2 := st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	st.TestSort("Quick Sort", qs.QuickSort, arr2, n)

}
