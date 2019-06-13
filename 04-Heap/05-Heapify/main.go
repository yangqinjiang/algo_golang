package main

import (
	core3 "algo_golang/02-Sorting-Basic/core"
	core2 "algo_golang/03-Sorting-Advance/core"
	"algo_golang/04-Heap/core"
	"fmt"
)

func main() {

	n := 1000000
	ms := new(core2.MergeSort)
	qs_advance := new(core2.QuickSortAdvance)
	qs_2_ways := new(core2.QuickSortTwoWays)
	qs_3_ways := new(core2.QuickSortThreeWays)
	// 测试1 一般性测试
	fmt.Printf("Test for random array,size = %d , random range [0,%d]\n", n, n)

	st := new(core3.SortTestHelper)
	arr1 := st.GenerateRandomArray(n, 0, n)
	arr2 := st.CopyIntArray(arr1, n)
	arr3 := st.CopyIntArray(arr1, n)
	arr4 := st.CopyIntArray(arr1, n)
	arr5 := st.CopyIntArray(arr1, n)
	arr6 := st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr2, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr3, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr4, n)
	st.TestSort("Heap Sort 1", core.HeapSort1, arr5, n)
	st.TestSort("Heap Sort 2 Heapify", core.HeapSort2, arr6, n)

	fmt.Println("-------------")

	// 测试2 测试近乎有序的数组
	swapTimes := 100
	fmt.Printf("Test for nearly ordered array, size = %d , swap time = %d \n", n, swapTimes)

	arr1 = st.GenerateNearlyOrderedArray(n, swapTimes)
	arr2 = st.CopyIntArray(arr1, n)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)
	arr6 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr2, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr3, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr4, n)
	st.TestSort("Heap Sort 1", core.HeapSort1, arr5, n)
	st.TestSort("Heap Sort 2 Heapify", core.HeapSort2, arr6, n)

	fmt.Println("-------------")

	// 测试3 测试存在包含大量相同元素的数组
	fmt.Printf("Test for random array, size = %d , random range [0,10] \n", n)

	arr1 = st.GenerateRandomArray(n, 0, 10)
	//arr2 = st.CopyIntArray(arr1, n)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)
	arr6 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	// 这种情况下, 普通的QuickSort退化为O(n^2)的算法, 不做测试
	//st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr2, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr3, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr4, n)
	st.TestSort("Heap Sort 1", core.HeapSort1, arr5, n)
	st.TestSort("Heap Sort 2 Heapify", core.HeapSort2, arr6, n)

	fmt.Println("-------------")
}
