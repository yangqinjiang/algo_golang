package main

import (
	"fmt"
	core2 "algo_golang/03-Sorting-Advance/core"
	"algo_golang/02-Sorting-Basic/core"
)

func main() {

	fmt.Println("running Quick Sort advance")

	// 比较Merge Sort和Quick Sort两种排序算法的性能效率
	// 两种排序算法虽然都是O(nlogn)级别的, 但是Quick Sort算法有常数级的优势
	// Quick Sort要比Merge Sort快, 即使我们对Merge Sort进行了优化

	n := 1000000
	ms := new(core2.MergeSort)
	qs := new(core2.QuickSort)
	qs_advance := new(core2.QuickSortAdvance)
	// 测试1 一般性测试
	fmt.Printf("Test for random array,size = %d , random range [0,%d]\n", n, n)

	st := new(core.SortTestHelper)
	arr1 := st.GenerateRandomArray(n, 0, n)
	arr2 := st.CopyIntArray(arr1, n)
	arr3 := st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	st.TestSort("Quick Sort", qs.QuickSort, arr2, n)
	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr3, n)

	fmt.Println("-------------")
	// 测试2 测试近乎有序的数组
	// 加入了随机选择标定点的步骤后, 我们的快速排序可以轻松处理近乎有序的数组
	// 但是对于近乎有序的数组, 其效率比优化后的归并排序要低, 但完全再容忍范围里
	// 思考一下为什么对于近乎有序的数组, 快排的性能比优化后的归并排序低? :)
	swapTimes := 100
	fmt.Printf("Test for nearly ordered array, size = %d , swap time =  [0,%d]\n", n, swapTimes)
	arr1 = st.GenerateNearlyOrderedArray(n, swapTimes)
	arr2 = st.CopyIntArray(arr1, n)
	arr3 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	st.TestSort("Quick Sort", qs.QuickSort, arr2, n)
	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr3, n)

}
